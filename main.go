package main

import (
	"net/http"

	"github.com/levidurfee/ham/handlers"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	appengine.Main()
}
