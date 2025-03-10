package service

import (
	"context"
	"gomall/app/payment/biz/dal/mysql"
	payment "gomall/rpc_gen/kitex_gen/payment"
	"testing"

	"github.com/joho/godotenv"
)

func TestCharge_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		Amount: 100,
		OrderId: "1",
		UserId: 1,
        CreditCard: &payment.CreditCardInfo{
		CreditCardNumber: "4111111111111111",
		CreditCardCvv: 123,
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: 2030,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
