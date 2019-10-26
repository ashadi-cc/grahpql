package pg

import (
	"fmt"
	"gql-ashadi/model"
	"gql-ashadi/repository"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/xid"
)

type UserRepo struct{}

var userRepo *UserRepo
var userInit sync.Once

//GetUserRepo get instance repository
func GetUserRepo() repository.User {
	userInit.Do(func() {
		userRepo = &UserRepo{}
	})

	return userRepo
}

//GetByEmail get user by email
func (repo *UserRepo) GetByEmail(email string) (*model.User, error) {
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

//Create user
func (repo *UserRepo) Create(user *model.User) error {
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

//GetCount get total record
func (repo *UserRepo) GetCount() (int, error) {
	var totalRecord int
	sqlStatement := "SELECT COUNT(id) AS total FROM USERS"
	smt, err := GetInstance().Prepare(sqlStatement)
	if err != nil {
		return totalRecord, errors.Wrapf(err, "problem when preparing query :%s", sqlStatement)
	}
	defer smt.Close()

	if err := smt.QueryRow().Scan(&totalRecord); err != nil {
		return totalRecord, errors.Wrap(err, "problem scan row")
	}
	return totalRecord, nil
}

//GetUsers get user with limit and page
func (repo *UserRepo) GetUsers(args model.PageInfo) (*model.Users, error) {
	if args.Page <= 0 || args.Limit <= 0 {
		return nil, fmt.Errorf("Limit and page could not 0 %+v", args)
	}
	sqlStatement := "SELECT id,email,first_name,last_name FROM users ORDER BY created_at LIMIT $1 OFFSET $2"
	smt, err := GetInstance().Prepare(sqlStatement)
	if err != nil {
		return nil, errors.Wrapf(err, "problem when preparing query :%s", sqlStatement)
	}
	defer smt.Close()

	var (
		totalCount int
		errTotal   error
		wgroup     sync.WaitGroup
	)
	wgroup.Add(1)
	go func(wg *sync.WaitGroup) {
		totalCount, errTotal = repo.GetCount()
		wg.Done()
	}(&wgroup)
	wgroup.Wait()

	if errTotal != nil {
		return nil, err
	}

	offset := (args.Limit * (args.Page - 1))
	rows, err := smt.Query(args.Limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "failed when query users")
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	var (
		UserID    string
		email     string
		firstName string
		lastName  string
	)
	for rows.Next() {
		if err := rows.Scan(&UserID, &email, &firstName, &lastName); err != nil {
			return nil, errors.Wrap(err, "error when scan row")
		}
		users = append(users, &model.User{
			ID:        UserID,
			Email:     email,
			FirstName: firstName,
			LastName:  lastName,
		})
	}

	UserInfo := model.Users{
		TotalCount: totalCount,
		Items:      users,
		PageInfo: &model.PageInfo{
			Page:  args.Page,
			Limit: args.Limit,
		},
	}
	return &UserInfo, nil
}
