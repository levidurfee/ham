package sess

import (
	"context"
	"net/http"

	sessions "github.com/gorilla/Sessions"
)

// Get a session variable
func Get(ctx context.Context, w http.ResponseWriter, r *http.Request, key string) (interface{}, error) {
	store := sessions.NewCookieStore(getSessionKey(ctx))

	session, err := store.Get(r, "ham-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return session.Values[key], nil
}

// Save a session variable
func Save(ctx context.Context, w http.ResponseWriter, r *http.Request, key string, value interface{}) {
	store := sessions.NewCookieStore(getSessionKey(ctx))

	session, _ := store.Get(r, "ham-session")

	session.Save(r, w)
}
