package service

import (
	"context"
	"testing"

	"gomall_study/app/cart/biz/dal/mysql"
	"gomall_study/app/cart/infra/rpc"
	cart "gomall_study/rpc_gen/kitex_gen/cart"

	"github.com/joho/godotenv"
)

func TestAddItem_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	rpc.InitClient()

	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 2,
			Quantity:  1,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
