package data

import (
	"context"
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"

	_ "github.com/hi20160616/hfcms-users/api/users/v1"
	_ "github.com/hi20160616/hfcms-users/configs"
	"github.com/hi20160616/hfcms-users/internal/biz"
	"github.com/hi20160616/hfcms-users/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.DepartmentRepo = new(departmentRepo)

type departmentRepo struct {
	data *Data
	log  *log.Logger
}

func NewDepartmentRepo(data *Data, logger *log.Logger) biz.DepartmentRepo {
	return &departmentRepo{
		data: data,
		log:  log.Default(),
	}
}

// parent=categories/*/departments
// TODO parent=tags/*/departments
// parent=departments/*/departments
func (ur *departmentRepo) ListDepartments(ctx context.Context, parent string) (*biz.Departments, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	bizas := &biz.Departments{}
	us := &mariadb.Departments{}
	var err error
	re := regexp.MustCompile(`^(departments|usergroups|users)/(.+)/departments$`)
	x := re.FindStringSubmatch(parent)
	y, err := regexp.MatchString(parent, `^departments$`)
	if err != nil {
		return nil, err
	}
	if len(x) != 3 && y {
		us, err = ur.data.DBClient.DatabaseClient.QueryDepartment().All(ctx)
	} else {
		clause := [4]string{}
		switch x[1] {
		case "departments":
			clause = [4]string{"department_id", "=", x[2], "and"}
		case "usergroups":
			clause = [4]string{"usergroup_id", "=", x[2], "and"}
		case "users":
			clause = [4]string{"user_id", "=", x[2], "and"}
		}
		us, err = ur.data.DBClient.DatabaseClient.QueryDepartment().
			Where(clause).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, u := range us.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Department{
			DepartmentId: u.DepartmentId,
			State:        u.State,
			Deleted:      u.Deleted,
			UpdateTime:   timestamppb.New(u.UpdateTime),
		})
	}
	return bizas, nil
}

func (ur *departmentRepo) GetDepartment(ctx context.Context, name string) (*biz.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// name=departments/1
	re := regexp.MustCompile(`^departments/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id, "and"}
	u, err := ur.data.DBClient.DatabaseClient.QueryDepartment().
		Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Department{
		DepartmentId: u.DepartmentId,
		State:        u.State,
		Deleted:      u.Deleted,
		UpdateTime:   timestamppb.New(u.UpdateTime),
	}, nil
}

func (ur *departmentRepo) CreateDepartment(ctx context.Context, department *biz.Department) (*biz.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := ur.data.DBClient.DatabaseClient.
		InsertDepartment(ctx, &mariadb.Department{}); err != nil {
		return nil, err
	}
	return department, nil
}

func (ur *departmentRepo) UpdateDepartment(ctx context.Context, department *biz.Department) (*biz.Department, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbDepartment, err := ur.data.DBClient.DatabaseClient.QueryDepartment().
		Where([4]string{"id", "=", strconv.Itoa(department.DepartmentId), "and"}).
		First(ctx)
	if err != nil {
		return nil, err
	}
	if &department.State != nil {
		dbDepartment.State = department.State
	}
	if err := ur.data.DBClient.DatabaseClient.
		UpdateDepartment(ctx, dbDepartment); err != nil {
		return nil, err
	}
	return &biz.Department{
		DepartmentId: dbDepartment.DepartmentId,
		State:        dbDepartment.State,
		Deleted:      dbDepartment.Deleted,
		UpdateTime:   timestamppb.New(dbDepartment.UpdateTime),
	}, nil
}

// DeleteDepartment is soft delete, that can be undeleted, it just update deleted field to 1
func (ur *departmentRepo) DeleteDepartment(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^departments/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("departmentRepo: DeleteDepartment: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("departmentRepo: DeleteDepartment: department id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.DeleteDepartment(ctx, id)
}

func (ur *departmentRepo) UndeleteDepartment(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	re := regexp.MustCompile(`^departments/([\d.]+)/undelete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("departmentRepo: DeleteDepartment: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("departmentRepo: DeleteDepartment: department id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.UndeleteDepartment(ctx, id)
}

// DeleteDepartment2 is true delete row from database permanently, be careful
func (ur *departmentRepo) DeleteDepartment2(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^departments/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("departmentRepo: DeleteDepartment: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("departmentRepo: DeleteDepartment: department id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.DeleteDepartment2(ctx, id)
}
