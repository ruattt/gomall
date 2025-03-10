package service

import (
	"context"
	// "log"

	auth "gomall/app/frontend/hertz_gen/frontend/auth"
	"gomall/app/frontend/infra/rpc"
	"gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {

	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}


	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	err = session.Save()
	if err != nil {
		return "", err
	}

	// session_id := sessions.DefaultMany(h.RequestContext,"user_id")
	// session_id.Set("user_id", resp.UserId)
	// session_email := sessions.DefaultMany(h.RequestContext,"user_email")
	// session_email.Set("email", req.Email)
	// err = session_id.Save()
	// if err != nil {
	// 	return "", err
	// }
	// err = session_email.Save()
	// if err != nil {
	// 	return "", err
	// }

	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
