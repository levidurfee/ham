package sess

import (
	"context"
	"net/http"

	sessions "github.com/gorilla/Sessions"
)

const defaultSession = "hamsession"

// Get a session variable
func Get(ctx context.Context, w http.ResponseWriter, r *http.Request, key string) (interface{}, error) {
	sk, err := getSessionKey(ctx)
	if err != nil {
		httpSessionError(w, err)
		return nil, err
	}

	store := sessions.NewCookieStore(sk)

	session, err := store.Get(r, defaultSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return session.Values[key], nil
}

// Save a session variable
func Save(ctx context.Context, w http.ResponseWriter, r *http.Request, key string, value interface{}) error {
	sk, err := getSessionKey(ctx)
	if err != nil {
		httpSessionError(w, err)
		return err
	}

	session, err := sessions.NewCookieStore(sk).Get(r, defaultSession)
	if err != nil {
		httpSessionError(w, err)
		return err
	}

	session.Save(r, w)

	return nil
}

func httpSessionError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
