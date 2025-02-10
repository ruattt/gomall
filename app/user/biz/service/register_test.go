package service

import (
	"context"
	"testing"

	"gomall_study/app/user/biz/dal/mysql"
	user "gomall_study/app/user/kitex_gen/user"

	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "1demo@damin.com",
		Password:        "FJODIAFUFJO",
		PasswordConfirm: "FJODIAFUFJO",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
