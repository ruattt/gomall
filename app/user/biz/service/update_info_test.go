package service

import (
	"context"
	"gomall/app/user/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"testing"

	"github.com/joho/godotenv"
)

func TestUpdateInfo_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewUpdateInfoService(ctx)
	// init req and assert value

	req := &user.UpdateInfoReq{
		UserId:           1,
        Email:          "abc@a.c",
        Password:        "",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
