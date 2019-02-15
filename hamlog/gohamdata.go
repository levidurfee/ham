package hamlog

import (
	"net/http"

	"github.com/levidurfee/ham/id"
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

// GOhamData is a struct for storing basic data
type GOhamData struct {
	LoggedIn  bool
	Login     string
	Logout    string
	Template  string
	User      *user.User
	RequestID int64
}

// BuildData from request for template
func BuildData(r *http.Request) GOhamData {
	ctx := appengine.NewContext(r)
	ctx = id.CtxWithID(ctx)
	u := user.Current(ctx)
	var g GOhamData
	g.LoggedIn = true
	g.User = u
	g.RequestID = id.GetID(ctx)
	if u == nil {
		g.LoggedIn = false
	}
	login, _ := user.LoginURL(ctx, "/")
	logout, _ := user.LogoutURL(ctx, "/")

	g.Login = login
	g.Logout = logout

	id.PrintID(ctx)

	return g
}
