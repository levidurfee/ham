package models

import (
	"net/http"

	"google.golang.org/appengine/log"

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
	if err != nil {
		log.Debugf(ctx, "No token cookie", nil)
	}
	if err == nil {
		g.Token = token.Value
		g.HAM, err = user.NewHAM(r)
		if err == nil {
			g.UserID = g.HAM.UID
			//log.Debugf(ctx, "%v", g.UserID)
		}
	}

	id.PrintID(ctx)

	return g
}
