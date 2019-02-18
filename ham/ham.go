package ham

import (
	"encoding/base64"
	"os"

	sessions "github.com/gorilla/Sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine"
)

var (
	// OAuthConfig contains the configuration options for Google oauth
	OAuthConfig *oauth2.Config

	// SessionStore has the gorilla
	SessionStore sessions.Store
)

func init() {

	// [START auth]
	OAuthConfig = configureOAuthClient(os.Getenv("OAUTH_CLIENT_ID"), os.Getenv("OAUTH_CLIENT_SECRET"))
	// [END auth]

	// [START Sessions]
	sessionKey, _ := base64.StdEncoding.DecodeString(os.Getenv("SESSION_SECRET_KEY"))
	sessionEncKey, _ := base64.StdEncoding.DecodeString(os.Getenv("SESSION_SECRET_KEY"))
	cookieStore := sessions.NewCookieStore(sessionKey, sessionEncKey)
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	SessionStore = cookieStore
	// [END sessions]
}

func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := os.Getenv("OAUTH2_CALLBACK")

	// Check if this is running on the dev server. If it is, change the oauth2
	// callback URL to localhost 8080
	if appengine.IsDevAppServer() {
		redirectURL = "http://localhost:8080/oauth2callback"
	}

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
