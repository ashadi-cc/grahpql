package resolver

import (
	"gql-ashadi/model"

	graphql "github.com/graph-gophers/graphql-go"
)

type UserResolver struct {
	user *model.User
}

func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}

func (r *UserResolver) Email() string {
	return r.user.Email
}

func (r *UserResolver) FirstName() string {
	return r.user.FirstName
}

func (r *UserResolver) LastName() string {
	return r.user.LastName
}

func NewUserResolver(user *model.User) *UserResolver {
	return &UserResolver{user: user}
}
