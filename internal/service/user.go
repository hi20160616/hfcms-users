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

type UserService struct {
	pb.UnimplementedUsersAPIServer
	uc *biz.UserUsecase
}

func NewUserService() (*UserService, error) {
	dbc, err := mariadb.NewClient("hfcms-users")
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewUserRepo(db, log.Default())
	userUsecase := biz.NewUserUsecase(repo, *log.Default())
	return &UserService{uc: userUsecase}, nil
}

func (as *UserService) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListUsers: \n%v\n", r)
		}
	}()
	bizus, err := as.uc.ListUsers(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.User{}
	for _, u := range bizus.Collection {
		resp = append(resp, &pb.User{
			UserId:     int32(u.UserId),
			Username:   u.Username,
			Password:   u.Password,
			Realname:   u.Realname,
			Nickname:   u.Nickname,
			AvatarUrl:  u.AvatarUrl,
			Phone:      u.Phone,
			UserIp:     u.UserIP,
			State:      int32(u.State),
			Deleted:    int32(u.Deleted),
			CreateTime: u.CreateTime,
			UpdateTime: u.UpdateTime,
		})
	}
	return &pb.ListUsersResponse{Users: resp}, nil
}

func (us *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetUser: %s\n%v\n", in.Name, r)
		}
	}()
	bizu, err := us.uc.GetUser(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		UserId:     int32(bizu.UserId),
		Username:   bizu.Username,
		Password:   bizu.Password,
		Realname:   bizu.Realname,
		Nickname:   bizu.Nickname,
		AvatarUrl:  bizu.AvatarUrl,
		Phone:      bizu.Phone,
		UserIp:     bizu.UserIP,
		State:      int32(bizu.State),
		Deleted:    int32(bizu.Deleted),
		CreateTime: bizu.CreateTime,
		UpdateTime: bizu.UpdateTime,
	}, nil
}

func (us *UserService) SearchUsers(ctx context.Context, in *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchUsers: \n%v\n", r)
		}
	}()
	bizus, err := us.uc.SearchUsers(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchUsersResponse{}
	for _, e := range bizus.Collection {
		resp.Users = append(resp.Users, &pb.User{
			UserId:     int32(e.UserId),
			Username:   e.Username,
			Password:   e.Password,
			Realname:   e.Realname,
			Nickname:   e.Nickname,
			AvatarUrl:  e.AvatarUrl,
			Phone:      e.Phone,
			UserIp:     e.UserIP,
			State:      int32(e.State),
			Deleted:    int32(e.Deleted),
			CreateTime: e.CreateTime,
			UpdateTime: e.UpdateTime,
		})
	}
	return resp, nil
}

func (as *UserService) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateUsers: \n%v\n", r)
		}
	}()
	return nil, nil
	// a, err := as.ac.UpdateUser(ctx, &biz.User{
	//         UserId:  in.User.UserId,
	//         Title:      in.User.Title,
	//         Content:    in.User.Content,
	//         CategoryId: int(in.User.CategoryId),
	//         UserId:     int(in.User.UserId),
	// })
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.User{
	//         UserId:  a.UserId,
	//         Title:      a.Title,
	//         Content:    a.Content,
	//         CategoryId: int32(a.CategoryId),
	//         UserId:     int32(a.UserId),
	//         UpdateTime: a.UpdateTime,
	// }, nil
}

func (as *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateUsers: \n%v\n", r)
		}
	}()
	return nil, nil
	// return as.ac.DeleteUser(ctx, in.Name)
}

func (as *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateUsers: \n%v\n", r)
		}
	}()
	return nil, nil
	// a, err := as.ac.CreateUser(ctx, &biz.User{
	//         UserId:  in.User.UserId,
	//         Title:      in.User.Title,
	//         Content:    in.User.Content,
	//         CategoryId: int(in.User.CategoryId),
	//         UserId:     int(in.User.UserId),
	// })
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.User{
	//         UserId:  a.UserId,
	//         Title:      a.Title,
	//         Content:    a.Content,
	//         CategoryId: int32(a.CategoryId),
	//         UserId:     int32(a.UserId),
	//         UpdateTime: a.UpdateTime,
	// }, nil
}
