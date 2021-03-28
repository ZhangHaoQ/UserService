package kit

import (
	pb "UserService/proto"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type UserServer struct {
	LoginHandler    grpctransport.Handler
	RegisterHandler grpctransport.Handler
}

func (this *UserServer) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {
	_, res, err := this.LoginHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.LoginRes), nil
}

func (this *UserServer) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	_, res, err := this.RegisterHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.RegisterRes), nil
}
