package repos

import (
	"net/http"

	"github.com/levidurfee/ham/sess"

	"github.com/levidurfee/ham/models"
	"google.golang.org/appengine"
)

// BuildApp constructs a struct with app data.
func BuildApp(w http.ResponseWriter, r *http.Request) models.App {
	// First we'll create an App struct. We go ahead and set the LoggedIn field
	// false as a default, we'll check cookies to see if they're logged in, and
	// can update this field if they are logged in.
	app := models.App{
		LoggedIn: false,
	}

	ctx := appengine.NewContext(r)
	loggedIn, _ := sess.Get(ctx, w, r, "loggedin")
	_ = loggedIn

	return app
}
