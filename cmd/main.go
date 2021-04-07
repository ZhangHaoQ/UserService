package main

import (
	"UserService/Domain/repository"
	"UserService/Domain/service"
	"UserService/Register"
	"UserService/config"
	pb "UserService/proto"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Panic(err)
	}
	err = repository.Init(cfg)
	if err != nil {
		log.Panic(err)
	}
	namingClient, err := Register.InitNacos()
	if err != nil {
		log.Panic(err)
	}

	_, err = Register.RegisterService(namingClient)
	if err != nil {
		log.Panic(err)
	}

	UserServe := service.UserServer{}

	UserLoginHandler := grpctransport.NewServer(service.MakeUserLoginEndPoint(), service.DecodeRequest, service.EncodeRequest)
	UserRegisterHandler := grpctransport.NewServer(service.MakeUserRegisterEndPoint(), service.DecodeRequest, service.EncodeRequest)

	UserServe.LoginHandler = UserLoginHandler
	UserServe.RegisterHandler = UserRegisterHandler

	serviceAddress := ":9600"
	ls, _ := net.Listen("tcp", serviceAddress)

	gs := grpc.NewServer()

	pb.RegisterUserServiceServer(gs, &UserServe)

	_ = gs.Serve(ls)

}
