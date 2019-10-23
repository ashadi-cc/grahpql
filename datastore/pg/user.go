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

var userRepo UserRepo
var userInit sync.Once

//GetUserRepo get instance repository
func GetUserRepo() repository.User {
	userInit.Do(func() {
		userRepo = UserRepo{}
	})

	return userRepo
}

//GetByEmail get user by email
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

//Create user
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

//GetUsers get user with limit and page
func (repo UserRepo) GetUsers(args model.PageInfo) (*model.Users, error) {
	if args.Page <= 0 || args.Limit <= 0 {
		return nil, fmt.Errorf("Limit and page could not 0 %+v", args)
	}
	sqlStatement := "SELECT id,email,first_name,last_name FROM users Order by created_at"
	sqlStatement = fmt.Sprintf(`WITH cte AS (%s) SELECT * FROM (TABLE cte LIMIT $1 OFFSET $2)
	SUB RIGHT JOIN (SELECT count(*) FROM cte) c(full_count) ON true
	WHERE $3 < full_count `, sqlStatement)

	smt, err := GetInstance().Prepare(sqlStatement)
	if err != nil {
		return nil, errors.Wrapf(err, "problem when preparing query :%s", sqlStatement)
	}
	defer smt.Close()

	offset := (args.Limit * (args.Page - 1))
	rows, err := smt.Query(args.Limit, offset, offset)
	if err != nil {
		return nil, errors.Wrap(err, "failed when query users")
	}

	users := make([]*model.User, 0)
	var (
		UserID     string
		email      string
		firstName  string
		lastName   string
		totalCount int
	)
	for rows.Next() {
		if err := rows.Scan(&UserID, &email, &firstName, &lastName, &totalCount); err != nil {
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
