package data

import (
	"context"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/hi20160616/hfcms-users/api/users/v1"
	_ "github.com/hi20160616/hfcms-users/configs"
	"github.com/hi20160616/hfcms-users/internal/biz"
	"github.com/hi20160616/hfcms-users/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.RoleRepo = new(roleRepo)

type roleRepo struct {
	data *Data
	log  *log.Logger
}

func NewRoleRepo(data *Data, logger *log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.Default(),
	}
}

// parent=categories/*/roles
// TODO parent=tags/*/roles
// parent=roles/*/roles
func (ur *roleRepo) ListRoles(ctx context.Context, parent string) (*biz.Roles, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	bizas := &biz.Roles{}
	us := &mariadb.Roles{}
	var err error
	re := regexp.MustCompile(`^(departments|roles|rolegroups)/(.+)/roles$`)
	x := re.FindStringSubmatch(parent)
	y, err := regexp.MatchString(parent, `^roles$`)
	if err != nil {
		return nil, err
	}
	if len(x) != 3 && y {
		us, err = ur.data.DBClient.DatabaseClient.QueryRole().All(ctx)
	} else {
		clause := [4]string{}
		switch x[1] {
		case "departments":
			clause = [4]string{"department_id", "=", x[2], "and"}
		case "roles":
			clause = [4]string{"role_id", "=", x[2], "and"}
		case "rolegroups":
			clause = [4]string{"rolegroup_id", "=", x[2], "and"}
		}
		us, err = ur.data.DBClient.DatabaseClient.QueryRole().
			Where(clause).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, u := range us.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Role{
			RoleId:     u.Id,
			State:      u.State,
			Deleted:    u.Deleted,
			UpdateTime: timestamppb.New(u.UpdateTime),
		})
	}
	return bizas, nil
}

func (ur *roleRepo) GetRole(ctx context.Context, name string) (*biz.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// name=roles/1
	re := regexp.MustCompile(`^roles/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	clause := [4]string{"id", "=", id, "and"}
	u, err := ur.data.DBClient.DatabaseClient.QueryRole().
		Where(clause).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Role{
		RoleId:     u.Id,
		State:      u.State,
		Deleted:    u.Deleted,
		UpdateTime: timestamppb.New(u.UpdateTime),
	}, nil
}

func (ur *roleRepo) SearchRoles(ctx context.Context, name string) (*biz.Roles, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^roles/(.+)/search$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	kws := strings.Split(
		strings.TrimSpace(strings.ReplaceAll(x[1], "　", " ")), ",")
	cs := [][4]string{}
	for _, kw := range kws {
		cs = append(cs,
			// cs will be filtered by Where(clauses...)
			// the last `or` `and` in clause will cut off.
			// so, every clause need `or` or `and` for last element.
			[4]string{"rolename", "like", kw, "or"},
			[4]string{"realname", "like", kw, "or"},
			[4]string{"nickname", "like", kw, "or"},
			[4]string{"role_ip", "like", kw, "and"},
		)
	}
	us, err := ur.data.DBClient.DatabaseClient.QueryRole().
		Where(cs...).All(ctx)
	if err != nil {
		return nil, err
	}
	bizas := &biz.Roles{Collection: []*biz.Role{}}
	for _, e := range us.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Role{
			RoleId:     e.Id,
			Rolename:   e.Rolename,
			Password:   e.Password,
			Realname:   e.Realname,
			Nickname:   e.Nickname,
			AvatarUrl:  e.AvatarUrl,
			Phone:      e.Phone,
			RoleIP:     e.RoleIP,
			State:      e.State,
			Deleted:    e.Deleted,
			CreateTime: timestamppb.New(e.CreateTime),
			UpdateTime: timestamppb.New(e.UpdateTime),
		})
	}
	return bizas, nil
}

func (ur *roleRepo) CreateRole(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := ur.data.DBClient.DatabaseClient.
		InsertRole(ctx, &mariadb.Role{
			Rolename:  role.Rolename,
			Password:  role.Password,
			Realname:  role.Realname,
			Nickname:  role.Nickname,
			AvatarUrl: role.AvatarUrl,
			Phone:     role.Phone,
			RoleIP:    role.RoleIP,
		}); err != nil {
		return nil, err
	}
	return role, nil
}

func (ur *roleRepo) UpdateRole(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	dbRole, err := ur.data.DBClient.DatabaseClient.QueryRole().
		Where([4]string{"id", "=", strconv.Itoa(role.RoleId), "and"}).
		First(ctx)
	if err != nil {
		return nil, err
	}
	if &role.State != nil {
		dbRole.State = role.State
	}
	if err := ur.data.DBClient.DatabaseClient.
		UpdateRole(ctx, dbRole); err != nil {
		return nil, err
	}
	return &biz.Role{
		RoleId:     dbRole.Id,
		State:      dbRole.State,
		Deleted:    dbRole.Deleted,
		UpdateTime: timestamppb.New(dbRole.UpdateTime),
	}, nil
}

// DeleteRole is soft delete, that can be undeleted, it just update deleted field to 1
func (ur *roleRepo) DeleteRole(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^roles/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("roleRepo: DeleteRole: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("roleRepo: DeleteRole: role id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.DeleteRole(ctx, id)
}

func (ur *roleRepo) UndeleteRole(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	re := regexp.MustCompile(`^roles/([\d.]+)/undelete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("roleRepo: DeleteRole: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("roleRepo: DeleteRole: role id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.UndeleteRole(ctx, id)
}

// DeleteRole2 is true delete row from database permanently, be careful
func (ur *roleRepo) DeleteRole2(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^roles/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{},
			errors.New("roleRepo: DeleteRole: name cannot match regex express")
	}
	id, err := strconv.Atoi(x[1])
	if err != nil {
		return nil,
			errors.New("roleRepo: DeleteRole: role id should be integer only")
	}
	return &emptypb.Empty{}, ur.data.DBClient.DatabaseClient.DeleteRole2(ctx, id)
}
