package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

// LogSheet entry
type LogSheet struct {
	Date        string
	CallSign    string
	RSTSent     int
	RSTReceived int
	Frequency   float64
	Mode        string
	Power       string
	QTH         string
	Country     string
	Comments    string
	Band        int
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handle)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3001", nil))
}
