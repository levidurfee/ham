package models

import (
	"net/http"

<<<<<<< HEAD
	"google.golang.org/appengine/log"

=======
>>>>>>> master
	"github.com/levidurfee/ham/id"
	"github.com/levidurfee/ham/user"
	"google.golang.org/appengine"
)

// PageData is a struct for storing basic data
type PageData struct {
	Title     string
	LoggedIn  bool
	Template  string
	RequestID int64
	Token     string
	UserID    string
	HAM       *user.HAM
}

// NewPageData is a construct for the PageData struct
func NewPageData(r *http.Request) PageData {
	ctx := appengine.NewContext(r)
	ctx = id.CtxWithID(ctx)
	var g PageData
	g.RequestID = id.GetID(ctx)
	g.UserID = ""
	token, err := r.Cookie("token")
<<<<<<< HEAD
	if err != nil {
		log.Debugf(ctx, "No token cookie", nil)
	}
	if err == nil {
		g.Token = token.Value
=======

	if err == nil {
		g.Token = token.Value

		// Store some of this in memcache so it doesn't have to hit the Firebase
		// API each pageload. There is a noticeable difference in speed on the
		// site when you're logged in vs. when you're not logged in. So using
		// memcache might help this.
>>>>>>> master
		g.HAM, err = user.NewHAM(r)
		if err == nil {
			g.UserID = g.HAM.UID
			//log.Debugf(ctx, "%v", g.UserID)
		}
	}

	id.PrintID(ctx)

	return g
}
