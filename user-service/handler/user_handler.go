package handler

import (
	"encoding/json"
	"net/http"
	"user-service/model"
	"user-service/service"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request){
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	service.CreateUser(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}