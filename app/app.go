package app

import (
	"net/http"
	"time"

	"github.com/levidurfee/ham/sess"

	"github.com/levidurfee/ham/user"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// App model contains an assortment of information
type App struct {
	LoggedIn bool
	User     user.User
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

	if app.LoggedIn {
		email, _ := sess.Get(ctx, w, r, "email")
		uid, _ := sess.Get(ctx, w, r, "uid")

		app.User.Email = email.(string)
		app.User.UID = uid.(string)
	}

	// Now we check to see if they have a token set. This will tell us if they
	// are trying to login during this request. If they're trying to login,
	// then we want to verify their token.
	if user.IsLoggingIn(w, r) && app.LoggedIn == false {
		// Check their token
		log.Debugf(ctx, "Logging in", app)
		tkn, _ := r.Cookie("token")

		u, _ := user.VerifyToken(w, r, tkn.Value)

		app.User = u
		expiration := time.Now().Add(time.Second)
		cookie := http.Cookie{Name: "token", Value: "", Expires: expiration}
		http.SetCookie(w, &cookie)

		if u.UID != "" {
			app.LoggedIn = true
			sess.Save(ctx, w, r, "loggedin", "true")
			sess.Save(ctx, w, r, "uid", u.UID)
			sess.Save(ctx, w, r, "email", u.Email)
		}

		return app
	}

	log.Debugf(ctx, "App [%v]", app)

	return app
}
