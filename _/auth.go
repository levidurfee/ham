package main

import (
	"context"
	"encoding/gob"
	"errors"
	"net/http"
	"net/url"

	"github.com/gofrs/uuid"
	"github.com/levidurfee/ham/ham"
	"golang.org/x/oauth2"
	plus "google.golang.org/api/plus/v1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const (
	defaultSessionID        = "default"
	googleProfileSessionKey = "google_profile"
	oauthTokenSessionKey    = "oauth_token"
	oauthFlowRedirectKey    = "redirect"
	loggedInSession         = "loggedin"
)

// Profile contains strings for user profile data
type Profile struct {
	ID, DisplayName, ImageURL string
}

// LoggedIn string type
type LoggedIn map[string]interface{}

func init() {
	// Gob encoding for gorilla/sessions
	gob.Register(&oauth2.Token{})
	gob.Register(&Profile{})
	gob.Register(&LoggedIn{})
}

// loginHandler initiates an OAuth flow to authenticate the user.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	sessionID := uuid.Must(uuid.NewV4()).String()

	oauthFlowSession, err := ham.SessionStore.New(r, sessionID)
	if err != nil {
		log.Debugf(ctx, "loginHandler error: [%v]", err)
	}
	oauthFlowSession.Options.MaxAge = 10 * 60 // 10 minutes

	redirectURL, err := validateRedirectURL(r.FormValue("redirect"))
	if err != nil {
		log.Debugf(ctx, "loginHandler error: [%v]", err)
	}
	oauthFlowSession.Values[oauthFlowRedirectKey] = redirectURL

	if err := oauthFlowSession.Save(r, w); err != nil {
		log.Debugf(ctx, "loginHandler error: [%v]", err)
	}

	// Use the session ID for the "state" parameter.
	// This protects against CSRF (cross-site request forgery).
	// See https://godoc.org/golang.org/x/oauth2#Config.AuthCodeURL for more detail.
	url := ham.OAuthConfig.AuthCodeURL(sessionID, oauth2.ApprovalForce,
		oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
}

// logoutHandler clears the default session.
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	session, err := ham.SessionStore.New(r, defaultSessionID)
	if err != nil {
		log.Errorf(ctx, "Logout error [%v]", err)
	}
	session.Options.MaxAge = -1 // Clear session.
	if err := session.Save(r, w); err != nil {
		log.Errorf(ctx, "Logout error [%v]", err)
	}
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/"
	}
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// validateRedirectURL checks that the URL provided is valid.
// If the URL is missing, redirect the user to the application's root.
// The URL must not be absolute (i.e., the URL must refer to a path within this
// application).
func validateRedirectURL(path string) (string, error) {
	if path == "" {
		return "/", nil
	}

	// Ensure redirect URL is valid and not pointing to a different server.
	parsedURL, err := url.Parse(path)
	if err != nil {
		return "/", err
	}
	if parsedURL.IsAbs() {
		return "/", errors.New("URL must not be absolute")
	}
	return path, nil
}

// oauthCallbackHandler completes the OAuth flow, retreives the user's profile
// information and stores it in a session.
func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	oauthFlowSession, err := ham.SessionStore.Get(r, r.FormValue("state"))
	if err != nil {
		log.Debugf(ctx, "oauthCallbackHandler error [oauthFlowSession] [%v]", err)
	}

	redirectURL, ok := oauthFlowSession.Values[oauthFlowRedirectKey].(string)
	// Validate this callback request came from the app.
	if !ok {
		log.Debugf(ctx, "oauthCallbackHandler error [redirectURL] [%v]", err)
	}

	code := r.FormValue("code")
	tok, err := ham.OAuthConfig.Exchange(ctx, code)
	if err != nil {
		log.Debugf(ctx, "oauthCallbackHandler error [tok] [%v]", err)
	}

	session, err := ham.SessionStore.New(r, defaultSessionID)
	if err != nil {
		log.Debugf(ctx, "oauthCallbackHandler error [session] [%v]", err)
	}

	profile, err := fetchProfile(ctx, tok)
	if err != nil {
		log.Debugf(ctx, "oauthCallbackHandler profile [%v]", err)
	}

	session.Values[oauthTokenSessionKey] = tok
	session.Values[loggedInSession] = "true"
	session.Values[googleProfileSessionKey] = stripProfile(profile)

	if err := session.Save(r, w); err != nil {
		log.Debugf(ctx, "oauthCallbackHandler error [%v]", err)
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func profileFromSession(r *http.Request) *Profile {
	session, err := ham.SessionStore.Get(r, defaultSessionID)
	if err != nil {
		return nil
	}
	tok, ok := session.Values[oauthTokenSessionKey].(*oauth2.Token)
	if !ok || !tok.Valid() {
		return nil
	}
	profile, ok := session.Values[googleProfileSessionKey].(*Profile)
	if !ok {
		return nil
	}
	return profile
}

func fetchProfile(ctx context.Context, tok *oauth2.Token) (*plus.Person, error) {
	client := oauth2.NewClient(ctx, ham.OAuthConfig.TokenSource(ctx, tok))
	plusService, err := plus.New(client)
	if err != nil {
		return nil, err
	}
	return plusService.People.Get("me").Do()
}

// stripProfile returns a subset of a plus.Person.
func stripProfile(p *plus.Person) *Profile {
	return &Profile{
		ID:          p.Id,
		DisplayName: p.DisplayName,
		ImageURL:    p.Image.Url,
	}
}

func isLoggedIn(r *http.Request) bool {
	ctx := appengine.NewContext(r)

	session, err := ham.SessionStore.Get(r, defaultSessionID)
	if err != nil {
		//log.Debugf(ctx, "Couldn't get session [%v]", err)
		return false
	}

	tok, ok := session.Values[oauthTokenSessionKey].(*oauth2.Token)
	if !ok || !tok.Valid() {
		//log.Debugf(ctx, "Couldn't get tok [%v]", ok)
		return false
	}

	loggedin, ok := session.Values[loggedInSession].(string)
	if !ok {
		log.Debugf(ctx, "Couldn't get loggedin: [%v]", loggedin)
		return false
	}

	return true
}
