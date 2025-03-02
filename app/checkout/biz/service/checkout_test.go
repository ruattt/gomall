package service

import (
	"context"
	"gomall_study/app/checkout/infra/rpc"
	checkout "gomall_study/rpc_gen/kitex_gen/checkout"
	"testing"

	"github.com/joho/godotenv"
)

func TestCheckout_Run(t *testing.T) {
	godotenv.Load("../../.env")

	rpc.InitClient()
	
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	// init req and assert value

	req := &checkout.CheckoutReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
