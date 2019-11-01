package user

import (
	"gql-ashadi/datastore/pg"
	"gql-ashadi/repository"
	"sync"
)

var userRepo repository.User
var userInit sync.Once

//GetService get user service
func GetService() repository.User {
	userInit.Do(func() {
		userRepo = pg.NewUser()
	})
	return userRepo
}
