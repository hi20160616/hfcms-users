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
	ac *biz.UserUsecase
}

func NewUserService() (*UserService, error) {
	dbc, err := mariadb.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewUserRepo(db, log.Default())
	userUsecase := biz.NewUserUsecase(repo, *log.Default())
	return &UserService{ac: userUsecase}, nil
}

func (as *UserService) ListUsers(ctx context.Context, in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListUsers: \n%v\n", r)
		}
	}()
	return nil, nil
	// bizas, err := as.ac.ListUsers(ctx, in.Parent)
	// if err != nil {
	//         return nil, err
	// }
	// resp := []*pb.User{}
	// for _, a := range bizas.Collection {
	//         resp = append(resp, &pb.User{
	//                 UserId:  a.UserId,
	//                 Title:      a.Title,
	//                 Content:    a.Content,
	//                 CategoryId: int32(a.CategoryId),
	//                 UserId:     int32(a.UserId),
	//                 Category:   getCate(a),
	//                 UpdateTime: a.UpdateTime,
	//         })
	// }
	// return &pb.ListUsersResponse{Users: resp}, nil
}

func (as *UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetUser: %s\n%v\n", in.Name, r)
		}
	}()
	return nil, nil
	// biza, err := as.ac.GetUser(ctx, in.Name)
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.User{
	//         UserId:  biza.UserId,
	//         Title:      biza.Title,
	//         Content:    biza.Content,
	//         CategoryId: int32(biza.CategoryId),
	//         UserId:     int32(biza.UserId),
	//         Category:   getCate(biza),
	//         Tags:       getTags(biza),
	//         Attributes: getAttrs(biza),
	//         UpdateTime: biza.UpdateTime,
	// }, nil
}

func (as *UserService) SearchUsers(ctx context.Context, in *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchUsers: \n%v\n", r)
		}
	}()
	return nil, nil
	// bizas, err := as.ac.SearchUsers(ctx, in.Name)
	// if err != nil {
	//         return nil, err
	// }
	// respAs := &pb.SearchUsersResponse{}
	// for _, a := range bizas.Collection {
	//         respAs.Users = append(respAs.Users, &pb.User{
	//                 UserId:  a.UserId,
	//                 Title:      a.Title,
	//                 Content:    a.Content,
	//                 CategoryId: int32(a.CategoryId),
	//                 UserId:     int32(a.UserId),
	//                 Category:   getCate(a),
	//                 UpdateTime: a.UpdateTime})
	// }
	// return respAs, nil
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
