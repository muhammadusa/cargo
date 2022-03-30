package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main () {

	r := mux.NewRouter()

	r.HandleFunc("/", PostClassifiedsCtrl).Methods("POST")
	r.HandleFunc("/", GetClassifiedsCtrl).Methods("GET")
	// r.HandleFunc("/", DelateClassifiedsCtrl).Methods("DELATE")

	http.ListenAndServe(":8080", r)

}