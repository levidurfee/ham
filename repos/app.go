package repos

import "github.com/levidurfee/ham/models"

// BuildApp constructs a struct with app data.
func BuildApp() *models.App {
	// First we'll create an App struct. We go ahead and set the LoggedIn field
	// false as a default, we'll check cookies to see if they're logged in, and
	// can update this field if they are logged in.
	app := &models.App{
		LoggedIn: false,
	}

	return app
}
