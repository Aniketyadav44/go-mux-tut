package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UsersData)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	paramId, _ := strconv.Atoi(vars["id"])

	for _, user := range UsersData {
		if user.ID == paramId {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

type ResponseMsg struct {
	Msg  string `json:"message"`
	User User   `json:"user"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")

	var user User
	user.ID = rand.Intn(10000)

	json.NewDecoder(r.Body).Decode(&user)
	UsersData = append(UsersData, user)

	msg := "User created successfully"
	json.NewEncoder(w).Encode(ResponseMsg{msg, user})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")

	vars := mux.Vars(r)
	paramId, _ := strconv.Atoi(vars["id"])

	for index, user := range UsersData {
		if user.ID == paramId {
			UsersData = append(UsersData[:index], UsersData[index+1:]...)
		}
	}

	var user User
	user.ID = paramId
	json.NewDecoder(r.Body).Decode(&user)
	UsersData = append(UsersData, user)

	msg := "User updated successfully"
	json.NewEncoder(w).Encode(ResponseMsg{msg, user})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")

	vars := mux.Vars(r)
	paramId, _ := strconv.Atoi(vars["id"])

	for index, user := range UsersData {
		if user.ID == paramId {
			UsersData = append(UsersData[:index], UsersData[index+1:]...)
		}
	}

	msg := "User deleted successfully"
	json.NewEncoder(w).Encode(msg)
}
