package pg

import (
	"gql-ashadi/model"
	"gql-ashadi/repository"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type UserRepo struct{}

var userRepo UserRepo
var userInit sync.Once

//GetUserRepo get instance repository
func GetUserRepo() repository.User {
	userInit.Do(func() {
		userRepo = UserRepo{}
	})

	return userRepo
}

func (repo UserRepo) GetByEmail(email string) (*model.User, error) {
	sqlStatement := "Select id,email,first_name,last_name From users Where email = $1"
	smt, err := GetInstance().Prepare(sqlStatement)
	if err != nil {
		return nil, errors.Wrapf(err, "problem when preparing query :%s", sqlStatement)
	}
	defer smt.Close()

	var user = model.User{}
	if err := smt.QueryRow(email).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName); err != nil {
		return nil, errors.Wrap(err, "problem when scan row")
	}

	return &user, nil
}

func (repo UserRepo) Create(user *model.User) error {
	user.ID = xid.New().String()
	sqlStatement := "Insert into users (id,email,first_name,last_name) Values($1,$2,$3,$4)"
	smt, err := GetInstance().Prepare(sqlStatement)
	if err != nil {
		return errors.Wrapf(err, "problem when preparing query :%s", sqlStatement)
	}
	defer smt.Close()

	if _, err := smt.Exec(user.ID, user.Email, user.FirstName, user.LastName); err != nil {
		return errors.Wrapf(err, "problem when inserting user")
	}

	return nil
}
