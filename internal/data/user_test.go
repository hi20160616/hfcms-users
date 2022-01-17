package data

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hi20160616/hfcms-users/internal/biz"
	"github.com/hi20160616/hfcms-users/internal/data/db/mariadb"
)

var ar = func() biz.UserRepo {
	dc, err := mariadb.NewClient("hfcms-users")
	if err != nil {
		log.Fatal(err)
	}
	return NewUserRepo(&Data{DBClient: dc}, log.Default())
}()

var id = "1"

func TestCreateUser(t *testing.T) {
	// a, err := ar.CreateUser(context.Background(), &biz.User{
	//         Title:      "Test Create user title",
	//         Content:    "Test Create user content",
	//         CategoryId: 1,
	//         UserId:     1,
	// })
	// if err != nil {
	//         t.Error(err)
	// }
	// fmt.Println(a.UserId)
}

func TestUpdateUser(t *testing.T) {
	// a, err := ar.UpdateUser(context.Background(), &biz.User{
	//         UserId:     id,
	//         Title:      "Test Update user title",
	//         Content:    "Test Update user content",
	//         CategoryId: 1,
	//         UserId:     1,
	// })
	// if err != nil {
	//         t.Error(err)
	// }
	// fmt.Println(a.UserId)
}

func TestGetUser(t *testing.T) {
	u, err := ar.GetUser(context.Background(), "users/"+id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func TestSearchUsers(t *testing.T) {
	out := func(name string) {
		fmt.Println("name: ", name)
		as, err := ar.SearchUsers(context.Background(), name)
		if err != nil {
			t.Error(err)
			return
		}
		for _, a := range as.Collection {
			fmt.Println(a)
		}
	}

	names := []string{
		"users/test1/search",
		"users/test1,test2/search",
	}
	for _, n := range names {
		out(n)
	}
}

func TestListUsers(t *testing.T) {
	as, err := ar.ListUsers(context.Background(), "")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Collection {
		fmt.Println(a)
	}
	as, err = ar.ListUsers(context.Background(), "users/")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Collection {
		fmt.Println(a)
	}
}

func TestDeleteUser(t *testing.T) {
	name := "users/" + id + "/delete"
	if _, err := ar.DeleteUser(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := ar.GetUser(context.Background(), "users/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a != nil {
		t.Error(fmt.Errorf("DeleteUser failed."))
	}
}
