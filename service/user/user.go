package user

import (
	"gql-ashadi/datastore/pg"
	"gql-ashadi/repository"
)

func GetService() repository.User {
	return pg.GetUserRepo()
}
