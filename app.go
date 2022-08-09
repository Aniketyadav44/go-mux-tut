package main

import (
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func HandleFunc() {
	Router.HandleFunc("/users", GetAllUsers).Methods("GET")
	Router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	Router.HandleFunc("/user", CreateUser).Methods("POST")
	Router.HandleFunc("/user/{id}", UpdateUser).Methods("PUT")
	Router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
}
