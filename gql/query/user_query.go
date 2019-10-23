package query

import (
	"fmt"
	"gql-ashadi/gql/resolver"
	"gql-ashadi/model"
	"gql-ashadi/service/logger"
	"gql-ashadi/service/user"
)

//User resolver user
func (r *Resolver) User(args struct{ Email string }) (*resolver.UserResolver, error) {
	user, err := user.GetService().GetByEmail(args.Email)
	var errResolver error
	if err != nil {
		errResolver = fmt.Errorf("email not found %s", args.Email)
		logger.GetLogger().Info("get user by email error", err.Error())
	}
	return resolver.NewUserResolver(user), errResolver
}

//Users resolver
func (r *Resolver) Users(args struct {
	Page  int32
	Limit int32
}) (*resolver.UsersResolver, error) {
	pageInfo := model.PageInfo{
		Page:  int(args.Page),
		Limit: int(args.Limit),
	}

	users, err := user.GetService().GetUsers(pageInfo)
	var errResolver error
	if err != nil {
		errResolver = fmt.Errorf("errors when get users")
		logger.GetLogger().Info("unable to get users", err.Error())
	}
	return resolver.NewUsersResolver(users), errResolver
}
