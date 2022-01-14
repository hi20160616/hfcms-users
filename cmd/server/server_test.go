package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "github.com/hi20160616/hfcms-users/api/users/v1"
	"github.com/hi20160616/hfcms-users/configs"
	"google.golang.org/grpc"
)

func TestGRPCServer(t *testing.T) {
	tt, err := time.ParseDuration("1s")
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), tt)
	defer cancel()

	cfg := configs.NewConfig("hfcms-users")
	// Set up a connection to the server
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	// users
	c := pb.NewUsersAPIClient(conn)
	as, err := c.ListUsers(ctx, &pb.ListUsersRequest{Parent: ""})
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range as.Users {
		fmt.Printf("%-5d %-30s %-30s \n", a.UserId, a.Nickname, a.Realname)
	}

	id := "1"
	a, err := c.GetUser(context.Background(), &pb.GetUserRequest{Name: "users/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)

	// categories
	// c := pb.NewCategoriesAPIClient(conn)
	// cs, err := c.ListCategories(ctx, &pb.ListCategoriesRequest{})
	// if err != nil {
	//         log.Fatal(err)
	// }
	// for _, c := range cs.Categories {
	//         fmt.Printf("%-5d %-30s %-30s \n", c.CategoryId, c.CategoryName, c.CategoryCode)
	// }
	//
	// id := "3"
	// cc, err := c.GetCategory(ctx, &pb.GetCategoryRequest{Name: "categories/" + id})
	// if err != nil {
	//         log.Fatal(err)
	// }
	// fmt.Println(cc)

	// tags
	// tc := pb.NewTagsAPIClient(conn)
	// ts, err := tc.ListTags(ctx, &pb.ListTagsRequest{})
	// if err != nil {
	//         log.Fatal(err)
	// }
	// for _, tag := range ts.Tags {
	//         fmt.Printf("%-5d %-30s %-30s \n", tag.TagId, tag.TagName, tag.UpdateTime)
	// }
}
