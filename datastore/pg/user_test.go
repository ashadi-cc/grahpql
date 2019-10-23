package pg

import (
	"gql-ashadi/model"
	"testing"
)

func TestGetUsers(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	args := model.PageInfo{
		Limit: 10,
		Page:  1,
	}

	_, err := GetUserRepo().GetUsers(args)
	if err != nil {
		t.Fatalf("error when get users: %s", err.Error())
	}
}
