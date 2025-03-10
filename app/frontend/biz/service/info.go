package service

import (
	"context"
	"log"

	auth "gomall/app/frontend/hertz_gen/frontend/auth"
	"gomall/app/frontend/infra/rpc"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type InfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewInfoService(Context context.Context, RequestContext *app.RequestContext) *InfoService {
	return &InfoService{RequestContext: RequestContext, Context: Context}
}

func (h *InfoService) Run(req *auth.UserInfoReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	userInfoResp, err := rpc.UserClient.GetInfo(h.Context, &user.GetInfoReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	log.Printf("get user info success: %v", userInfoResp)
	return utils.H{
		"userEmail": userInfoResp.Email,
		"UserName":  userInfoResp.Name,
	}, nil
}
