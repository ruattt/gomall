package user

import (
	"context"
	user "gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetInfo(ctx context.Context, req *user.GetInfoReq, callOptions ...callopt.Option) (resp *user.GetInfoResp, err error) {
	resp, err = defaultClient.GetInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateInfo(ctx context.Context, req *user.UpdateInfoReq, callOptions ...callopt.Option) (resp *user.UpdateInfoResp, err error) {
	resp, err = defaultClient.UpdateInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
