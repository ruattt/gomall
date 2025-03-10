package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	product "gomall/rpc_gen/kitex_gen/product"
	"testing"

	"github.com/joho/godotenv"
)

func TestListProducts_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
