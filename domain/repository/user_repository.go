package irepository

import "go-rest-api/domain/model"

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
