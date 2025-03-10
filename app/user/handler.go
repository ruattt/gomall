package main

import (
	"context"
	user "gomall/rpc_gen/kitex_gen/user"
	"gomall/app/user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// GetInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetInfo(ctx context.Context, req *user.GetInfoReq) (resp *user.GetInfoResp, err error) {
	resp, err = service.NewGetInfoService(ctx).Run(req)

	return resp, err
}

// UpdateInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateInfo(ctx context.Context, req *user.UpdateInfoReq) (resp *user.UpdateInfoResp, err error) {
	resp, err = service.NewUpdateInfoService(ctx).Run(req)

	return resp, err
}
