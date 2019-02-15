package user

import (
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	firebase "firebase.google.com/go"
)

var (
	firebaseConfig = &firebase.Config{
		DatabaseURL:   "https://hamradio-96e65.firebaseio.com",
		ProjectID:     "hamradio",
		StorageBucket: "hamradio.appspot.com",
	}
)

// HAM is a ham
type HAM struct {
	Email string
	UID   string
	In    bool
}

// NewHAM creates a new HAM
func NewHAM(r *http.Request) (*HAM, error) {
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

	// Get the Cookie "token", we hope it exists
	t, err := r.Cookie("token")
	if err != nil {
		log.Errorf(ctx, "Cookie Error %v", err)
	}

	// Validate the token we received
	tok, err := auth.VerifyIDTokenAndCheckRevoked(ctx, t.Value)
	if err != nil {
		log.Infof(ctx, "auth.VerifyIDAndCheckRevoked: %v", err)
		return nil, err
	}

	// Get the User from the token
	user, err := auth.GetUser(ctx, tok.UID)
	if err != nil {
		log.Errorf(ctx, "auth.GetUser: %v", err)
	}

	// Log the user information, for now
	log.Debugf(ctx, "%v", user.UID)

	ham := &HAM{
		Email: user.Email,
		UID:   user.UID,
		In:    true,
	}

	return ham, nil
}
