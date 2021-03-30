package main

import (
	"UserService/Dao/repository"
	"UserService/config"
	"UserService/kit"
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

	UserServe := kit.UserServer{}

	UserLoginHandler := grpctransport.NewServer(kit.MakeUserLoginEndPoint(), kit.DecodeRequest, kit.EncodeRequest)
	UserRegisterHandler := grpctransport.NewServer(kit.MakeUserRegisterEndPoint(), kit.DecodeRequest, kit.EncodeRequest)

	UserServe.LoginHandler = UserLoginHandler
	UserServe.RegisterHandler = UserRegisterHandler

	serviceAddress := ":9600"
	ls, _ := net.Listen("tcp", serviceAddress)

	gs := grpc.NewServer()

	pb.RegisterUserServiceServer(gs, &UserServe)

	_ = gs.Serve(ls)

}
