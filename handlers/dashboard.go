package handlers

import (
	"net/http"

	"github.com/levidurfee/ham/models"
)

// DashboardHandler will handle loading everything for the dashboard
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	p := models.NewPageData(w, r)
	p.Template = "dashboard.html"

	RenderTemplate(w, p)
}
