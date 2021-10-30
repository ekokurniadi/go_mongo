package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getAllUsers).Methods("GET")
	r.HandleFunc("/users/profile", getUserProfile).Methods("POST")
	r.HandleFunc("/users", updateProfile).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
