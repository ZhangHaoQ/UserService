package main

import (
	pb "UserService/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9600", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	Uclient := pb.NewUserServiceClient(conn)

	res, err := Uclient.Login(context.Background(), &pb.LoginReq{Username: "zhq", Pwd: "123456"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
	fmt.Println(res.Code)
}
