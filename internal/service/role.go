package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hi20160616/hfcms-users/api/users/v1"
	"github.com/hi20160616/hfcms-users/internal/biz"
	"github.com/hi20160616/hfcms-users/internal/data"
	"github.com/hi20160616/hfcms-users/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RoleService struct {
	pb.UnimplementedRolesAPIServer
	uc *biz.RoleUsecase
}

func NewRoleService() (*RoleService, error) {
	dbc, err := mariadb.NewClient("hfcms-users")
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewRoleRepo(db, log.Default())
	userUsecase := biz.NewRoleUsecase(repo, *log.Default())
	return &RoleService{uc: userUsecase}, nil
}

func (as *RoleService) ListRoles(ctx context.Context, in *pb.ListRolesRequest) (*pb.ListRolesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListRoles: \n%v\n", r)
		}
	}()
	bizus, err := as.uc.ListRoles(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Role{}
	for _, u := range bizus.Collection {
		resp = append(resp, &pb.Role{
			RoleId:      int32(u.RoleId),
			RoleName:    u.RoleName,
			RoleCode:    u.RoleCode,
			Description: u.Description,
			State:       int32(u.State),
			Deleted:     int32(u.Deleted),
			UpdateTime:  u.UpdateTime,
		})
	}
	return &pb.ListRolesResponse{Roles: resp}, nil
}

func (us *RoleService) GetRole(ctx context.Context, in *pb.GetRoleRequest) (*pb.Role, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetRole: %s\n%v\n", in.Name, r)
		}
	}()
	bizu, err := us.uc.GetRole(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Role{
		RoleId:     int32(bizu.RoleId),
		State:      int32(bizu.State),
		Deleted:    int32(bizu.Deleted),
		UpdateTime: bizu.UpdateTime,
	}, nil
}

func (us *RoleService) SearchRoles(ctx context.Context, in *pb.SearchRolesRequest) (*pb.SearchRolesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchRoles: \n%v\n", r)
		}
	}()
	bizus, err := us.uc.SearchRoles(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchRolesResponse{}
	for _, e := range bizus.Collection {
		resp.Roles = append(resp.Roles, &pb.Role{
			RoleId:     int32(e.RoleId),
			State:      int32(e.State),
			Deleted:    int32(e.Deleted),
			UpdateTime: e.UpdateTime,
		})
	}
	return resp, nil
}

func (as *RoleService) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest) (*pb.Role, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateRoles: \n%v\n", r)
		}
	}()
	return nil, nil
	// a, err := as.ac.UpdateRole(ctx, &biz.Role{
	//         RoleId:  in.Role.RoleId,
	//         Title:      in.Role.Title,
	//         Content:    in.Role.Content,
	//         CategoryId: int(in.Role.CategoryId),
	//         RoleId:     int(in.Role.RoleId),
	// })
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.Role{
	//         RoleId:  a.RoleId,
	//         Title:      a.Title,
	//         Content:    a.Content,
	//         CategoryId: int32(a.CategoryId),
	//         RoleId:     int32(a.RoleId),
	//         UpdateTime: a.UpdateTime,
	// }, nil
}

func (as *RoleService) DeleteRole(ctx context.Context, in *pb.DeleteRoleRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateRoles: \n%v\n", r)
		}
	}()
	return nil, nil
	// return as.ac.DeleteRole(ctx, in.Name)
}

func (as *RoleService) CreateRole(ctx context.Context, in *pb.CreateRoleRequest) (*pb.Role, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateRoles: \n%v\n", r)
		}
	}()
	return nil, nil
	// a, err := as.ac.CreateRole(ctx, &biz.Role{
	//         RoleId:  in.Role.RoleId,
	//         Title:      in.Role.Title,
	//         Content:    in.Role.Content,
	//         CategoryId: int(in.Role.CategoryId),
	//         RoleId:     int(in.Role.RoleId),
	// })
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.Role{
	//         RoleId:  a.RoleId,
	//         Title:      a.Title,
	//         Content:    a.Content,
	//         CategoryId: int32(a.CategoryId),
	//         RoleId:     int32(a.RoleId),
	//         UpdateTime: a.UpdateTime,
	// }, nil
}
