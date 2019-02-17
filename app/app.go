package app

import (
	"net/http"

	"github.com/levidurfee/ham/user"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// App model contains an assortment of information
type App struct {
	LoggedIn bool
}

// NewApp constructs a struct with app data.
func NewApp(w http.ResponseWriter, r *http.Request) App {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "Creating new App", nil)

	// First we'll create an App struct. We go ahead and set the LoggedIn field
	// false as a default, we'll check cookies to see if they're logged in, and
	// can update this field if they are logged in.
	app := App{
		LoggedIn: false,
	}

	// I'll want to see if there is a loggedin cookie set.
	app.LoggedIn = user.LoggedIn(w, r)

	// Now we check to see if they have a token set. This will tell us if they
	// are trying to login during this request. If they're trying to login,
	// then we want to verify their token.
	if user.IsLoggingIn(w, r) {
		// Check their token
		log.Debugf(ctx, "Logging in", app)
	}

	return app
}
