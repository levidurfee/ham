package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/user/create/", createUserHandler)

	appengine.Main()
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Gopher Network!")
}
