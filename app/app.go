package app

import (
	"net/http"

	"github.com/levidurfee/ham/sess"

	"google.golang.org/appengine"
)

// App model contains an assortment of information
type App struct {
	LoggedIn bool
}

// NewApp constructs a struct with app data.
func NewApp(w http.ResponseWriter, r *http.Request) App {
	// First we'll create an App struct. We go ahead and set the LoggedIn field
	// false as a default, we'll check cookies to see if they're logged in, and
	// can update this field if they are logged in.
	app := App{
		LoggedIn: false,
	}

	ctx := appengine.NewContext(r)
	loggedIn, _ := sess.Get(ctx, w, r, "loggedin")
	_ = loggedIn

	return app
}

// func (h *App) VerifyToken(token string) (bool, error) {

// }
