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

var repo = func() biz.UserRepo {
	dc, err := mariadb.NewClient("hfcms-users")
	if err != nil {
		log.Fatal(err)
	}
	return NewUserRepo(&Data{DBClient: dc}, log.Default())
}()

var id = "9"

func TestCreateUser(t *testing.T) {
	_, err := repo.CreateUser(context.Background(), &biz.User{
		Username:  "dataInserted",
		Password:  "123",
		Realname:  "aaa",
		Nickname:  "bbb",
		AvatarUrl: "ccc.jpg",
		Phone:     "13412345678",
		UserIP:    "123.123.123.123",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUser(t *testing.T) {
	_, err := repo.UpdateUser(context.Background(), &biz.User{
		UserId:   9,
		Username: "dataUpdated",
		Password: "dataUpdated",
		Realname: "dataUpdated",
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetUser(t *testing.T) {
	u, err := repo.GetUser(context.Background(), "users/"+id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func TestSearchUsers(t *testing.T) {
	out := func(name string) {
		fmt.Println("name: ", name)
		es, err := repo.SearchUsers(context.Background(), name)
		if err != nil {
			t.Error(err)
			return
		}
		for _, a := range es.Collection {
			fmt.Println(a)
		}
	}

	names := []string{
		"users/test1/search",
		"users/test,lisi/search",
	}
	for _, n := range names {
		out(n)
	}
}

func TestListUsers(t *testing.T) {
	es, err := repo.ListUsers(context.Background(), "")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range es.Collection {
		fmt.Println(a)
	}
	es, err = repo.ListUsers(context.Background(), "users/")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range es.Collection {
		fmt.Println(a)
	}
}

func TestDeleteUser(t *testing.T) {
	name := "users/" + id + "/delete"
	if _, err := repo.DeleteUser(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := repo.GetUser(context.Background(), "users/"+id)
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

func TestUndeleteUser(t *testing.T) {
	name := "users/" + id + "/undelete"
	if _, err := repo.UndeleteUser(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := repo.GetUser(context.Background(), "users/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a.Deleted == 1 {
		t.Error(fmt.Errorf("UndeleteUser failed."))
	}

}
