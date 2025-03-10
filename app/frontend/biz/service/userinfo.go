package service

import (
	"context"
	
	auth "gomall/app/frontend/hertz_gen/frontend/auth"
	"gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type UserinfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserinfoService(Context context.Context, RequestContext *app.RequestContext) *UserinfoService {
	return &UserinfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UserinfoService) Run(req *auth.UserInfoReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	userInfoResp, err := rpc.UserClient.GetInfo(h.Context, &user.GetInfoReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"UserEmail": userInfoResp.Email,
		"UserName":  userInfoResp.Name,
	}, nil
}
