package service

import (
	"context"
	"gomall/app/user/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetInfo_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()

	ctx := context.Background()
	s := NewGetInfoService(ctx)
	// init req and assert value

	req := &user.GetInfoReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
