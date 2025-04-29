package service

import "tatar/book/.gen/mydb/model"

type UserService interface {
	GetUsers(limit int, offset int) ([]*model.User, error)
	GetUserByUUID(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(id string) error
}
