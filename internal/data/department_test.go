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

var dr = func() biz.DepartmentRepo {
	dc, err := mariadb.NewClient("hfcms-users")
	if err != nil {
		log.Fatal(err)
	}
	return NewDepartmentRepo(&Data{DBClient: dc}, log.Default())
}()

var departmentId = "9"

func TestCreateDepartment(t *testing.T) {
	_, err := dr.CreateDepartment(context.Background(), &biz.Department{})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateDepartment(t *testing.T) {
	_, err := dr.UpdateDepartment(context.Background(), &biz.Department{
		DepartmentId: 9,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetDepartment(t *testing.T) {
	u, err := dr.GetDepartment(context.Background(), "departments/"+id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}

func TestListDepartments(t *testing.T) {
	es, err := dr.ListDepartments(context.Background(), "")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range es.Collection {
		fmt.Println(a)
	}
	es, err = dr.ListDepartments(context.Background(), "departments/")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range es.Collection {
		fmt.Println(a)
	}
}

func TestDeleteDepartment(t *testing.T) {
	name := "departments/" + id + "/delete"
	if _, err := dr.DeleteDepartment(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := dr.GetDepartment(context.Background(), "departments/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a != nil {
		t.Error(fmt.Errorf("DeleteDepartment failed."))
	}
}

func TestUndeleteDepartment(t *testing.T) {
	name := "departments/" + id + "/undelete"
	if _, err := dr.UndeleteDepartment(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := dr.GetDepartment(context.Background(), "departments/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a.Deleted == 1 {
		t.Error(fmt.Errorf("UndeleteDepartment failed."))
	}

}
