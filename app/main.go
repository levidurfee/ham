package main

import (
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", registerHandler)

	appengine.Main()
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom"))
}
