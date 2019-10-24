package pg

import (
	"fmt"
	"gql-ashadi/model"
	"log"
	"sync"
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

func BenchmarkCreatUser(b *testing.B) {
	var waitgroup sync.WaitGroup

	for i := 1; i < 100000; i++ {
		waitgroup.Add(1)
		go func(i int, w *sync.WaitGroup) {
			u := &model.User{
				Email:     fmt.Sprintf("c%d@gmail.com", i),
				FirstName: fmt.Sprintf("c%d@gmail.com", i),
				LastName:  fmt.Sprintf("c%d@gmail.com", i),
			}
			err := GetUserRepo().Create(u)
			defer w.Done()
			if err != nil {
				log.Fatal(err)
			}
		}(i, &waitgroup)
	}

	waitgroup.Wait()
}
