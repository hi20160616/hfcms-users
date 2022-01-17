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

var rr = func() biz.RoleRepo {
	dc, err := mariadb.NewClient("hfcms-users")
	if err != nil {
		log.Fatal(err)
	}
	return NewRoleRepo(&Data{DBClient: dc}, log.Default())
}()

var roleId = "9"

func TestCreateRole(t *testing.T) {
	_, err := rr.CreateRole(context.Background(), &biz.Role{})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateRole(t *testing.T) {
	_, err := rr.UpdateRole(context.Background(), &biz.Role{
		RoleId: 9,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetRole(t *testing.T) {
	u, err := rr.GetRole(context.Background(), "roles/"+id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func TestSearchRoles(t *testing.T) {
	out := func(name string) {
		fmt.Println("name: ", name)
		es, err := rr.SearchRoles(context.Background(), name)
		if err != nil {
			t.Error(err)
			return
		}
		for _, a := range es.Collection {
			fmt.Println(a)
		}
	}

	names := []string{
		"roles/test1/search",
		"roles/test,lisi/search",
	}
	for _, n := range names {
		out(n)
	}
}

func TestListRoles(t *testing.T) {
	es, err := rr.ListRoles(context.Background(), "")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range es.Collection {
		fmt.Println(a)
	}
	es, err = rr.ListRoles(context.Background(), "roles/")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range es.Collection {
		fmt.Println(a)
	}
}

func TestDeleteRole(t *testing.T) {
	name := "roles/" + id + "/delete"
	if _, err := rr.DeleteRole(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := rr.GetRole(context.Background(), "roles/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a != nil {
		t.Error(fmt.Errorf("DeleteRole failed."))
	}
}

func TestUndeleteRole(t *testing.T) {
	name := "roles/" + id + "/undelete"
	if _, err := rr.UndeleteRole(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := rr.GetRole(context.Background(), "roles/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a.Deleted == 1 {
		t.Error(fmt.Errorf("UndeleteRole failed."))
	}

}
