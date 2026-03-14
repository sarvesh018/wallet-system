package repository

import "user-service/model"

var users = make(map[string]model.User)

func SaveUser(user model.User){
	users[user.ID] = user
}

func GetUser(id string) (model.User, bool){
	user, exists := users[id]
	return user, exists
}
