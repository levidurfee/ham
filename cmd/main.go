package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)

	http.Handle("/", r)

	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
