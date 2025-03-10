package service

import (
	"context"
	"gomall/app/cart/biz/dal/mysql"
	cart "gomall/rpc_gen/kitex_gen/cart"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetCart_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewGetCartService(ctx)
	// init req and assert value

	req := &cart.GetCartReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
