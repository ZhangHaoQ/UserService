package service

import (
	"UserService/Domain/model"
	"UserService/Domain/repository"
	pb "UserService/proto"
	"context"
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/go-kit/kit/endpoint"
	xerrors "github.com/pkg/errors"
)

func MakeUserLoginEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LoginReq)
		username := req.Username
		pwd := req.Pwd
		UR := repository.NewUserDB(context.Background())
		if !UR.ExistUserByName(username) {
			response = &pb.RegisterRes{
				Code: 400,
				Msg:  "User not exist",
			}
			return response, nil
		}
		user, err := UR.GetPassWordByName(username)
		if err != nil {
			response = &pb.LoginRes{Code: 400, Msg: "NO!!!!"}
			return response, err
		}
		if pwd != user.Password {
			response = &pb.LoginRes{Code: 400, Msg: "PassWord Wrong!"}
			return response, nil
		}
		response = &pb.LoginRes{Code: 200, Msg: "okok!"}
		return response, nil
	}
}

func MakeUserRegisterEndPoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.RegisterReq)
		valid := validation.Validation{}
		valid.Required(req.Username, "user_name")
		valid.Required(req.Pwd, "password")
		valid.Phone(req.Mobile, "mobile")
		if valid.HasErrors() {
			response = &pb.RegisterRes{
				Code: 400,
				Msg:  "Data Error",
			}
			err = errors.New("DataError")
			return response, xerrors.Wrapf(err, "Validation error")
		}
		username := req.Username
		pwd := req.Pwd
		UR := repository.NewUserDB(context.Background())
		if UR.ExistUserByName(username) {
			response = &pb.RegisterRes{
				Code: 400,
				Msg:  "User already exist",
			}
			return response, nil
		}
		M := model.UserModel{
			UserName: username,
			Password: pwd,
			Mobile:   req.Mobile,
		}
		err = UR.Create(M)
		if err != nil {
			response = &pb.RegisterRes{
				Code: 500,
				Msg:  "Create Model Fail",
			}
			return response, err
		}
		response = &pb.RegisterRes{
			Code: 200,
			Msg:  "ok",
		}
		return response, nil
	}
}

func DecodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func EncodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
