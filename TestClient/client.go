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

	res, err := Uclient.Register(context.Background(), &pb.RegisterReq{Username: "LYY", Pwd: "741741741", Mobile: "17645051903"})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
	fmt.Println(res.Code)
}
