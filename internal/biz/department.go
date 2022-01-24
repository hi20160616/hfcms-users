package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Department struct {
	DepartmentId, ParentId, State, Deleted      int
	DepartmentCode, DepartmentName, Description string
	UpdateTime                                  *timestamppb.Timestamp
}

type Departments struct {
	Collection    []*Department
	NextPageToken string
}

type DepartmentRepo interface {
	ListDepartments(ctx context.Context, parent string) (*Departments, error)
	GetDepartment(ctx context.Context, name string) (*Department, error)
	CreateDepartment(ctx context.Context, article *Department) (*Department, error)
	UpdateDepartment(ctx context.Context, article *Department) (*Department, error)
	DeleteDepartment(ctx context.Context, name string) (*emptypb.Empty, error)
	UndeleteDepartment(ctx context.Context, name string) (*emptypb.Empty, error)
	DeleteDepartment2(ctx context.Context, name string) (*emptypb.Empty, error)
}

type DepartmentUsecase struct {
	repo DepartmentRepo
}

func NewDepartmentUsecase(repo DepartmentRepo, logger log.Logger) *DepartmentUsecase {
	return &DepartmentUsecase{repo: repo}
}

func (au *DepartmentUsecase) CreateDepartment(ctx context.Context, article *Department) (*Department, error) {
	return au.repo.CreateDepartment(ctx, article)
}

func (au *DepartmentUsecase) ListDepartments(ctx context.Context, parent string) (*Departments, error) {
	return au.repo.ListDepartments(ctx, parent)
}

func (au *DepartmentUsecase) GetDepartment(ctx context.Context, name string) (*Department, error) {
	return au.repo.GetDepartment(ctx, name)
}

func (au *DepartmentUsecase) UpdateDepartment(ctx context.Context, article *Department) (*Department, error) {
	return au.repo.UpdateDepartment(ctx, article)
}

func (au *DepartmentUsecase) DeleteDepartment(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.DeleteDepartment(ctx, name)
}
func (au *DepartmentUsecase) UndeleteDepartment(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.UndeleteDepartment(ctx, name)
}
func (au *DepartmentUsecase) DeleteDepartment2(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.DeleteDepartment2(ctx, name)
}
