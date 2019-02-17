package sess

import (
	"context"
	"net/http"
)

// Get a session variable
func Get(ctx context.Context, w http.ResponseWriter, r *http.Request, key string) (interface{}, error) {
	session, err := getSession(ctx, w, r)
	if err != nil {
		httpSessionError(w, err)
		return nil, err
	}

	return session.Values[key], nil
}
