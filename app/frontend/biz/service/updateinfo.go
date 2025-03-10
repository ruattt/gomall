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

type UpdateinfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateinfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateinfoService {
	return &UpdateinfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateinfoService) Run(req *auth.UpdateInfoReq) (resp map[string]any, err error) {
	// Finish your business logic.
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	UpdateInfoResp, err := rpc.UserClient.UpdateInfo(h.Context, &user.UpdateInfoReq{
		UserId:   uint32(userId),
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
        return nil, err
    }
	return utils.H{
		"Success": UpdateInfoResp.Success,
		"UserName":    UpdateInfoResp.Name,
		"UserEmail":   UpdateInfoResp.Email,
	}, nil
}
