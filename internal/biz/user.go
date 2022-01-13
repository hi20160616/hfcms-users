package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	UserId, State, Deleted                 int
	Username, Password, Realname, Nickname string
	AvatarUrl, Phone, UserIP               string
	int
	CreateTime, UpdateTime *timestamppb.Timestamp
}

type Users struct {
	Collection    []*User
	NextPageToken string
}

type UserRepo interface {
	ListUsers(ctx context.Context, parent string) (*Users, error)
	GetUser(ctx context.Context, name string) (*User, error)
	SearchUsers(ctx context.Context, name string) (*Users, error)
	CreateUser(ctx context.Context, article *User) (*User, error)
	UpdateUser(ctx context.Context, article *User) (*User, error)
	DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (au *UserUsecase) CreateUser(ctx context.Context, article *User) (*User, error) {
	return au.repo.CreateUser(ctx, article)
}

func (au *UserUsecase) ListUsers(ctx context.Context, parent string) (*Users, error) {
	return au.repo.ListUsers(ctx, parent)
}

func (au *UserUsecase) GetUser(ctx context.Context, name string) (*User, error) {
	return au.repo.GetUser(ctx, name)
}

func (au *UserUsecase) SearchUsers(ctx context.Context, name string) (*Users, error) {
	return au.repo.SearchUsers(ctx, name)
}

func (au *UserUsecase) UpdateUser(ctx context.Context, article *User) (*User, error) {
	return au.repo.UpdateUser(ctx, article)
}

func (au *UserUsecase) DeleteUser(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.DeleteUser(ctx, name)
}
