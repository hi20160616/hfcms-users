package mariadb

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var roleId = 5

func TestListRoles(t *testing.T) {
	c, err := NewClient("hfcms-roles")
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryRole().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertRole(t *testing.T) {
	c, err := NewClient("hfcms-roles")
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*Role{
		{
			Rolename:  "testInsert1",
			Password:  "testInsert1",
			Realname:  "Mazzy1",
			Nickname:  "donkey1",
			AvatarUrl: "testInsert1.jpg",
			Phone:     "13512345678",
			RoleIP:    "123.123.123.123",
		},
		{
			Rolename:  "testInsert2",
			Password:  "testInsert2",
			Realname:  "Mazzy2",
			Nickname:  "donkey2",
			AvatarUrl: "testInsert2.jpg",
			Phone:     "13512345678",
			RoleIP:    "123.123.123.123",
		},
		{
			Rolename:  "testInsert3",
			Password:  "testInsert3",
			Realname:  "Mazzy3",
			Nickname:  "donkey3",
			AvatarUrl: "testInsert3.jpg",
			Phone:     "13512345678",
			RoleIP:    "123.123.123.123",
		},
	}
	for _, tc := range tcs {
		err := c.DatabaseClient.InsertRole(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUpdateRole(t *testing.T) {
	c, err := NewClient("hfcms-roles")
	if err != nil {
		t.Fatal(err)
	}
	role := &Role{
		Id:        id,
		Rolename:  "tttest",
		Password:  "testNewPwd",
		Realname:  "real test",
		Nickname:  "nick test",
		AvatarUrl: "avatar_url.test.jpg",
		Phone:     "13512345678",
		RoleIP:    "111.111.111.111",
		State:     1,
	}
	getRole := func() *Role {
		ps := [4]string{"id", "=", strconv.Itoa(role.Id), "or"}
		got, err := c.DatabaseClient.QueryRole().Where(ps).First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getRole()
	if err := c.DatabaseClient.UpdateRole(context.Background(), role); err != nil {
		t.Error(err)
		return
	}
	after := getRole()
	if before.Password != after.Password {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.Password, after.Password))
		}
	}
	if before.Realname != after.Realname {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.Realname, after.Realname))
		}
	}
	if before.Nickname != after.Nickname {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.Nickname, after.Nickname))
		}
	}
	if before.AvatarUrl != after.AvatarUrl {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.AvatarUrl, after.AvatarUrl))
		}
	}
	if before.Phone != after.Phone {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.Phone, after.Phone))
		}
	}
	if before.RoleIP != after.RoleIP {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.RoleIP, after.RoleIP))
		}
	}
	if before.State != after.State {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				role.State, after.State))
		}
	}
}

func TestDeleteRole(t *testing.T) {
	c, err := NewClient("hfcms-roles")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeleteRole(context.Background(), id); err != nil {
		t.Fatalf("DeleteRole err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryRole().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryRole err: %v", err)
	}
	if got.Deleted != 1 {
		t.Error(errors.New("Delete failed."))
	}
}

func TestUnDeleteRole(t *testing.T) {
	c, err := NewClient("hfcms-roles")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.UndeleteRole(context.Background(), id); err != nil {
		t.Fatalf("UndeleteRole err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryRole().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryRole err: %v", err)
	}
	if got.Deleted != 0 {
		t.Error(errors.New("UnDelete failed."))
	}

}
