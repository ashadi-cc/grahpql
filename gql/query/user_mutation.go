package query

import (
	"fmt"
	"gql-ashadi/gql/resolver"
	"gql-ashadi/model"
	"gql-ashadi/service/logger"
	"gql-ashadi/service/user"
)

//CreateUser create user resolver
func (r *Resolver) CreateUser(args struct {
	Email     string
	FirstName string
	LastName  string
}) (*resolver.UserResolver, error) {
	logger.GetLogger().Info("user created")
	userModel := &model.User{
		Email:     args.Email,
		FirstName: args.FirstName,
		LastName:  args.LastName,
	}
	var errResolver error
	if err := user.GetService().Create(userModel); err != nil {
		logger.GetLogger().Warning("unable to create user", err.Error())
		errResolver = fmt.Errorf("failed created user, duplicate email")
	}
	return resolver.NewUserResolver(userModel), errResolver
}
