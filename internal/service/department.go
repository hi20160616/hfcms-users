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

type DepartmentService struct {
	pb.UnimplementedDepartmentsAPIServer
	uc *biz.DepartmentUsecase
}

func NewDepartmentService() (*DepartmentService, error) {
	dbc, err := mariadb.NewClient("hfcms-users")
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewDepartmentRepo(db, log.Default())
	userUsecase := biz.NewDepartmentUsecase(repo, *log.Default())
	return &DepartmentService{uc: userUsecase}, nil
}

func (as *DepartmentService) ListDepartments(ctx context.Context, in *pb.ListDepartmentsRequest) (*pb.ListDepartmentsResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListDepartments: \n%v\n", r)
		}
	}()
	bizus, err := as.uc.ListDepartments(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Department{}
	for _, u := range bizus.Collection {
		resp = append(resp, &pb.Department{
			DepartmentId:   int32(u.DepartmentId),
			DepartmentName: u.DepartmentName,
			Description:    u.Description,
			State:          int32(u.State),
			Deleted:        int32(u.Deleted),
			UpdateTime:     u.UpdateTime,
		})
	}
	return &pb.ListDepartmentsResponse{Departments: resp}, nil
}

func (us *DepartmentService) GetDepartment(ctx context.Context, in *pb.GetDepartmentRequest) (*pb.Department, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetDepartment: %s\n%v\n", in.Name, r)
		}
	}()
	bizu, err := us.uc.GetDepartment(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Department{
		DepartmentId: int32(bizu.DepartmentId),
		State:        int32(bizu.State),
		Deleted:      int32(bizu.Deleted),
		UpdateTime:   bizu.UpdateTime,
	}, nil
}

func (as *DepartmentService) UpdateDepartment(ctx context.Context, in *pb.UpdateDepartmentRequest) (*pb.Department, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateDepartments: \n%v\n", r)
		}
	}()
	return nil, nil
	// a, err := as.ac.UpdateDepartment(ctx, &biz.Department{
	//         DepartmentId:  in.Department.DepartmentId,
	//         Title:      in.Department.Title,
	//         Content:    in.Department.Content,
	//         CategoryId: int(in.Department.CategoryId),
	//         DepartmentId:     int(in.Department.DepartmentId),
	// })
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.Department{
	//         DepartmentId:  a.DepartmentId,
	//         Title:      a.Title,
	//         Content:    a.Content,
	//         CategoryId: int32(a.CategoryId),
	//         DepartmentId:     int32(a.DepartmentId),
	//         UpdateTime: a.UpdateTime,
	// }, nil
}

func (as *DepartmentService) DeleteDepartment(ctx context.Context, in *pb.DeleteDepartmentRequest) (*emptypb.Empty, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateDepartments: \n%v\n", r)
		}
	}()
	return nil, nil
	// return as.ac.DeleteDepartment(ctx, in.Name)
}

func (as *DepartmentService) CreateDepartment(ctx context.Context, in *pb.CreateDepartmentRequest) (*pb.Department, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateDepartments: \n%v\n", r)
		}
	}()
	return nil, nil
	// a, err := as.ac.CreateDepartment(ctx, &biz.Department{
	//         DepartmentId:  in.Department.DepartmentId,
	//         Title:      in.Department.Title,
	//         Content:    in.Department.Content,
	//         CategoryId: int(in.Department.CategoryId),
	//         DepartmentId:     int(in.Department.DepartmentId),
	// })
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.Department{
	//         DepartmentId:  a.DepartmentId,
	//         Title:      a.Title,
	//         Content:    a.Content,
	//         CategoryId: int32(a.CategoryId),
	//         DepartmentId:     int32(a.DepartmentId),
	//         UpdateTime: a.UpdateTime,
	// }, nil
}
