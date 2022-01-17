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
	// c := pb.NewUsersAPIClient(conn)
	// as, err := c.ListUsers(ctx, &pb.ListUsersRequest{Parent: "users"})
	// if err != nil {
	//         t.Fatal(err)
	// }
	// for _, a := range as.Users {
	//         fmt.Printf("%-5d %-10s %-10s \n", a.UserId, a.Nickname, a.Realname)
	// }

	c := pb.NewUsersAPIClient(conn)
	as, err := c.ListUsers(ctx, &pb.ListUsersRequest{Parent: ""})
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range as.Users {
		fmt.Printf("%-3d %-20s %-20s \n", a.UserId, a.Nickname, a.Realname)
	}
	// id := "1"
	// a, err := c.GetUser(context.Background(), &pb.GetUserRequest{Name: "users/" + id})
	// if err != nil {
	//         t.Fatal(err)
	// }
	// fmt.Println(a)
}
