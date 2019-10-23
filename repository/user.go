package repository

import "gql-ashadi/model"

//UserRepository user repository interface
type User interface {
	GetByEmail(email string) (*model.User, error)
	GetUsers(args model.PageInfo) (*model.Users, error)
	Create(user *model.User) error
}
