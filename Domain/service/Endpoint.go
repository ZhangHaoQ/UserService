package service

import (
	pb "UserService/proto"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

func MakeUserLoginEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LoginReq)
		username := req.Username
		pwd := req.Pwd
		fmt.Println("UserName: ", username)
		fmt.Println("PassWord: ", pwd)
		response = &pb.LoginRes{
			Code: 200,
			Msg:  "ok",
		}
		return response, nil
	}
}

func MakeUserRegisterEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.RegisterReq)
		username := req.Username
		pwd := req.Pwd
		fmt.Println("UserName: ", username)
		fmt.Println("PassWord: ", pwd)
		response = &pb.RegisterRes{Code: 200, Msg: "okok!"}
		return response, nil
	}
}

func DecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func EncodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
