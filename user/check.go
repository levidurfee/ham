package user

import (
	"net/http"

	"github.com/levidurfee/ham/sess"
	"google.golang.org/appengine"
)

// LoggedIn checks if there is a user logged in, trying to login, or if they are a guest.
func LoggedIn(w http.ResponseWriter, r *http.Request) bool {
	ctx := appengine.NewContext(r)
	loggedIn, err := sess.Get(ctx, w, r, "loggedin")
	if err != nil {
		return false
	}

	if loggedIn == nil {
		return false
	}

	return true
}

// IsLoggingIn checks if they user is trying to login
func IsLoggingIn(w http.ResponseWriter, r *http.Request) bool {
	// Check for token, if no token, they haven't tried to login
	token, err := r.Cookie("token")
	if err != nil {
		return false
	}

	// Token was empty, return false, they haven't tried to login
	if token.Value == "" {
		return false
	}

	return true
}
