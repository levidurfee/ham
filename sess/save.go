package sess

import (
	"context"
	"net/http"
)

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
