package mariadb

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var id = 5

func TestPrepareQuery(t *testing.T) {
	qc := &UserQuery{query: "SELECT * FROM users"}
	qc.Where(
		[4]string{"name", "like", "test", "and"},
		[4]string{"name", "like", "test1", "and"},
		[4]string{"name", "like", "test2", "and"},
		[4]string{"name", "like", "test3", "and"},
	)
	if err := qc.prepareQuery(context.Background()); err != nil {
		t.Error(err)
	}
	fmt.Println(qc.query, qc.args)
}

func TestListUsers(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryUser().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestWhereUsers(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}

	out := func(a [4]string) {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", a)
		got, err := c.DatabaseClient.QueryUser().Where(a).All(context.Background())
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		for _, e := range got.Collection {
			fmt.Println(e)
		}
		fmt.Println("===========================================")
	}

	outs := func(ps [][4]string) {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", ps)
		got, err := c.DatabaseClient.QueryUser().Where(ps...).All(context.Background())
		if err != nil {
			t.Error(err)
			return
		}
		// fmt.Println(got.Collection)
		for _, e := range got.Collection {
			fmt.Println(e)
		}
		fmt.Println("===========================================")
	}

	ps1 := [][4]string{
		{"title", "=", "test2 title", "and"},
		{"content", "like", "2", "and"},
		{"user_id", "=", "0"},
	}
	ps2 := [][4]string{
		{"title", "=", "test2 title", "or"},
		{"content", "like", "2", "or"},
		{"user_id", "=", "0"},
	}
	out(ps1[2]) // test one clause
	outs(ps1)
	outs(ps2)
}

func TestInsertUser(t *testing.T) {
	// c, err := NewClient("hfcms-users")
	// if err != nil {
	//         t.Fatal(err)
	// }
	// user1 := &User{
	//         Id:         time.Now().Format("060102150405.000000") + "00001",
	//         Title:      "test1 title",
	//         Content:    "test1 content",
	//         UserId:     1,
	//         CategoryId: 1,
	//         UpdateTime: time.Now(),
	// }
	// user2 := &User{
	//         Id:         time.Now().Format("060102150405.000000") + "00002",
	//         Title:      "test2 title",
	//         Content:    "test2 content",
	//         UserId:     2,
	//         CategoryId: 2,
	//         UpdateTime: time.Now(),
	// }
	// user3 := &User{
	//         Id:         time.Now().Format("060102150405.000000") + "00003",
	//         Title:      "test3 title",
	//         Content:    "test3 content",
	//         UserId:     3,
	//         CategoryId: 3,
	//         UpdateTime: time.Now(),
	// }
	// if err := c.DatabaseClient.InsertUser(context.Background(), user1); err != nil {
	//         t.Error(err)
	// }
	// if err := c.DatabaseClient.InsertUser(context.Background(), user2); err != nil {
	//         t.Error(err)
	// }
	// if err := c.DatabaseClient.InsertUser(context.Background(), user3); err != nil {
	//         t.Error(err)
	// }
}
func TestUpdateUser(t *testing.T) {
	// c, err := NewClient("hfcms-users")
	// if err != nil {
	//         t.Fatal(err)
	// }
	// user := &User{
	//         Id:         id,
	//         Title:      "Test title update",
	//         Content:    "Test content update",
	//         CategoryId: 5,
	//         UserId:     2,
	// }
	// if err := c.DatabaseClient.UpdateUser(context.Background(), user); err != nil {
	//         t.Error(err)
	//         return
	// }
	// ps := [4]string{"id", "=", user.Id}
	// got, err := c.DatabaseClient.QueryUser().Where(ps).First(context.Background())
	// if err != nil {
	//         t.Error(err)
	//         return
	// }
	// fmt.Println(got)
}

func TestDeleteUser(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeleteUser(context.Background(), id); err != nil {
		t.Fatalf("DeleteUser err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryUser().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryUser err: %v", err)
	}
	if got.Deleted != 1 {
		t.Error(errors.New("Delete failed."))
	}
}
