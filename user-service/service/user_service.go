package service

import (
	"user-service/model"
	"user-service/repository"
)

func CreateUser(user model.User){
	repository.SaveUser(user)
}

func GetUser(id string) (model.User, bool){
	return repository.GetUser(id)
}