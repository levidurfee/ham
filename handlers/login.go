package handlers

import (
	"net/http"

	"github.com/levidurfee/ham/models"
)

// LoginHandler handles loading the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	p := models.NewPageData(w, r)
	p.Template = "login.html"

}
