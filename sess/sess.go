package sess

import (
	"context"
	"net/http"

	sessions "github.com/gorilla/Sessions"
)

const defaultSession = "hamsession"

// Get a session variable
func Get(ctx context.Context, w http.ResponseWriter, r *http.Request, key string) (interface{}, error) {
	session, err := getSession(ctx, w, r)
	if err != nil {
		httpSessionError(w, err)
		return nil, err
	}

	return session.Values[key], nil
}

// Save a session variable
func Save(ctx context.Context, w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	session, err := getSession(ctx, w, r)
	if err != nil {
		httpSessionError(w, err)
		return err
	}

	session.Values[key] = value
	session.Save(r, w)

	return nil
}

func getSession(ctx context.Context, w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	sk, err := getSessionKey(ctx)
	if err != nil {
		httpSessionError(w, err)
		return nil, err
	}

	session, err := sessions.NewCookieStore(sk).Get(r, defaultSession)
	if err != nil {
		httpSessionError(w, err)
		return nil, err
	}

	return session, nil
}
