package sess

import (
	"context"
	"net/http"

	sessions "github.com/gorilla/Sessions"
)

const defaultSession = "hamsession"

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
