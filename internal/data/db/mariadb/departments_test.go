package mariadb

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

var departmentId = 5

func TestListDepartments(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryDepartment().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestInsertDepartment(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	tcs := []*Department{
		{
			DepartmentName: "zeng",
			DepartmentCode: "01",
			Description:    "keyitianjian",
		},
	}
	for _, tc := range tcs {
		err := c.DatabaseClient.InsertDepartment(context.Background(), tc)
		if err != nil {
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUpdateDepartment(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	department := &Department{
		DepartmentId:   departmentId,
		ParentId:       111,
		DepartmentCode: "updatecodetest",
		DepartmentName: "updatenametest",
		Description:    "updatedesctest",
		State:          1,
	}
	getDepartment := func() *Department {
		ps := [4]string{"id", "=", strconv.Itoa(department.DepartmentId), "or"}
		got, err := c.DatabaseClient.QueryDepartment().Where(ps).First(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		return got
	}

	before := getDepartment()
	if err := c.DatabaseClient.UpdateDepartment(context.Background(), department); err != nil {
		t.Error(err)
		return
	}
	after := getDepartment()
	if before.DepartmentCode != after.DepartmentCode {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				department.DepartmentCode, after.DepartmentCode))
		}
	}
	if before.DepartmentName != after.DepartmentName {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				department.DepartmentName, after.DepartmentName))
		}
	}
	if before.Description != after.Description {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %s, got: %s",
				department.Description, after.Description))
		}
	}
	if before.ParentId != after.ParentId {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				department.ParentId, after.ParentId))
		}
	}
	if before.State != after.State {
		if err != nil {
			t.Fatal(fmt.Errorf("want: %d, got: %d",
				department.State, after.State))
		}
	}
}

func TestDeleteDepartment(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.DeleteDepartment(context.Background(), id); err != nil {
		t.Fatalf("DeleteDepartment err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryDepartment().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryDepartment err: %v", err)
	}
	if got.Deleted != 1 {
		t.Error(errors.New("Delete failed."))
	}
}

func TestUndeleteDepartment(t *testing.T) {
	c, err := NewClient("hfcms-users")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.DatabaseClient.UndeleteDepartment(context.Background(), id); err != nil {
		t.Fatalf("UndeleteDepartment err: %v", err)
	}

	ps := [4]string{"id", "=", strconv.Itoa(id), "and"}
	got, err := c.DatabaseClient.QueryDepartment().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Fatalf("QueryDepartment err: %v", err)
	}
	if got.Deleted != 0 {
		t.Error(errors.New("UnDelete failed."))
	}

}
