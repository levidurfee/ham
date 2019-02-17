package user

import (
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var (
	firebaseConfig = &firebase.Config{
		DatabaseURL:   "https://hamradio-96e65.firebaseio.com",
		ProjectID:     "hamradio",
		StorageBucket: "hamradio.appspot.com",
	}
)

// User type
type User struct {
	Email string
	UID   string
}

// VerifyToken verifies the token
func VerifyToken(w http.ResponseWriter, r *http.Request, token string) (User, error) {
	var u User
	// Get a new Context from App Engine
	ctx := appengine.NewContext(r)

	opt := option.WithCredentialsFile("sak.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Debugf(ctx, "error initializing app: %v", err)
	}

	// Get a new Auth from the Firebase App
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Errorf(ctx, "app.Auth: %v", err)
	}

	// Validate the token we received
	tok, err := auth.VerifyIDTokenAndCheckRevoked(ctx, token)
	if err != nil {
		log.Infof(ctx, "auth.VerifyIDAndCheckRevoked: %v", err)
		return u, err
	}

	// Get the User from the token
	user, err := auth.GetUser(ctx, tok.UID)
	if err != nil {
		log.Errorf(ctx, "auth.GetUser: %v", err)
	}

	// Log the user information, for now
	log.Debugf(ctx, "%v", user.UID)

	u.Email = user.Email
	u.UID = user.UID

	return u, nil
}
