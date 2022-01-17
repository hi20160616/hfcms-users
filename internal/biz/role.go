package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Role struct {
	RoleId, ParentId, State, Deleted int
	RoleCode, RoleName, Description  string
	UpdateTime                       *timestamppb.Timestamp
}

type Roles struct {
	Collection    []*Role
	NextPageToken string
}

type RoleRepo interface {
	ListRoles(ctx context.Context, parent string) (*Roles, error)
	GetRole(ctx context.Context, name string) (*Role, error)
	SearchRoles(ctx context.Context, name string) (*Roles, error)
	CreateRole(ctx context.Context, article *Role) (*Role, error)
	UpdateRole(ctx context.Context, article *Role) (*Role, error)
	DeleteRole(ctx context.Context, name string) (*emptypb.Empty, error)
	UndeleteRole(ctx context.Context, name string) (*emptypb.Empty, error)
	DeleteRole2(ctx context.Context, name string) (*emptypb.Empty, error)
}

type RoleUsecase struct {
	repo RoleRepo
}

func NewRoleUsecase(repo RoleRepo, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{repo: repo}
}

func (au *RoleUsecase) CreateRole(ctx context.Context, article *Role) (*Role, error) {
	return au.repo.CreateRole(ctx, article)
}

func (au *RoleUsecase) ListRoles(ctx context.Context, parent string) (*Roles, error) {
	return au.repo.ListRoles(ctx, parent)
}

func (au *RoleUsecase) GetRole(ctx context.Context, name string) (*Role, error) {
	return au.repo.GetRole(ctx, name)
}

func (au *RoleUsecase) SearchRoles(ctx context.Context, name string) (*Roles, error) {
	return au.repo.SearchRoles(ctx, name)
}

func (au *RoleUsecase) UpdateRole(ctx context.Context, article *Role) (*Role, error) {
	return au.repo.UpdateRole(ctx, article)
}

func (au *RoleUsecase) DeleteRole(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.DeleteRole(ctx, name)
}
func (au *RoleUsecase) UndeleteRole(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.UndeleteRole(ctx, name)
}
func (au *RoleUsecase) DeleteRole2(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.DeleteRole2(ctx, name)
}
