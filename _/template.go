package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// AppTemplate is a template builder
type AppTemplate struct {
	t *template.Template
}

// PageData has the information needed for the page
type PageData struct {
	Data     interface{}
	LoggedIn bool
}

// Execute handles creating the template with the user info
func (t *AppTemplate) Execute(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := t.t.Execute(w, data); err != nil {
		ctx := appengine.NewContext(r)
		log.Errorf(ctx, "Could not create template: %v", err)
		return err
	}

	return nil
}

// ParseTemplate returns a new AppTemplate
func ParseTemplate(filename string) *AppTemplate {
	tmpl := template.Must(template.ParseFiles("templates/base.html"))

	// Put the named file into a template called "body"
	path := filepath.Join("templates", filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("could not read template: %v", err))
	}
	template.Must(tmpl.New("body").Parse(string(b)))

	return &AppTemplate{tmpl.Lookup("base.html")}
}
