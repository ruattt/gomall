package service

import (
	"context"
	"gomall_study/app/order/biz/dal/mysql"
	order "gomall_study/rpc_gen/kitex_gen/order"
	"testing"

	"github.com/joho/godotenv"
)

func TestPlaceOrder_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	req := &order.PlaceOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
