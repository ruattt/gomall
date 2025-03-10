package service

import (
	"context"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"
	user "gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetInfoService struct {
	ctx context.Context
} // NewGetInfoService new GetInfoService
func NewGetInfoService(ctx context.Context) *GetInfoService {
	return &GetInfoService{ctx: ctx}
}

// Run create note info
func (s *GetInfoService) Run(req *user.GetInfoReq) (resp *user.GetInfoResp, err error) {
	// Finish your business logic.
	userInfo, err := model.GetByUserID(mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(2004003, err.Error())
	}
	return &user.GetInfoResp{
		Email: userInfo.Email,
		Name:  userInfo.Name,
	}, nil
}
