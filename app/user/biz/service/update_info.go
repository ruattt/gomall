package service

import (
	"context"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"
	user "gomall/rpc_gen/kitex_gen/user"

	"golang.org/x/crypto/bcrypt"
)

type UpdateInfoService struct {
	ctx context.Context
} // NewUpdateInfoService new UpdateInfoService
func NewUpdateInfoService(ctx context.Context) *UpdateInfoService {
	return &UpdateInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateInfoService) Run(req *user.UpdateInfoReq) (resp *user.UpdateInfoResp, err error) {
	// Finish your business logic.
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(password),
		Name:           req.Name,
	}
	userinfo, err := model.Update(s.ctx, mysql.DB, newUser, req.UserId)
	if err != nil {
		return nil, err
	}

	return &user.UpdateInfoResp{
		Success: true,
		Name:    userinfo.Name,
		Email:   userinfo.Email,
	}, nil
}
