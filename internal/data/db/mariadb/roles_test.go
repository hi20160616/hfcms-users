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
	c, err := NewClient("hfcms-users")
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
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*Role{
		{
			RoleName:    "zeng",
			RoleCode:    "01",
			Description: "keyitianjian",
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
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	role := &Role{
		RoleId:      roleId,
		ParentId:    111,
		RoleCode:    "updatecodetest",
		RoleName:    "updatenametest",
		Description: "updatedesctest",
		State:       1,
	}
	getRole := func() *Role {
		ps := [4]string{"id", "=", strconv.Itoa(role.RoleId), "or"}
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
	if before.RoleCode != after.RoleCode {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.RoleCode, after.RoleCode))
		}
	}
	if before.RoleName != after.RoleName {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.RoleName, after.RoleName))
		}
	}
	if before.Description != after.Description {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				role.Description, after.Description))
		}
	}
	if before.ParentId != after.ParentId {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				role.ParentId, after.ParentId))
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
	c, err := NewClient("hfcms-users")
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

func TestUndeleteRole(t *testing.T) {
	c, err := NewClient("hfcms-users")
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
