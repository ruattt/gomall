package service

import (
	"context"
	"gomall_study/app/cart/biz/dal/mysql"
	cart "gomall_study/rpc_gen/kitex_gen/cart"
	"testing"

	"github.com/joho/godotenv"
)

func TestEmptyCart_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	// init req and assert value

	req := &cart.EmptyCartReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
