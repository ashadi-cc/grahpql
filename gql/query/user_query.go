package query

import (
	"fmt"
	"gql-ashadi/gql/resolver"
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
