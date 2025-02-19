package service

import (
	"context"

	common "gomall_study/app/frontend/hertz_gen/frontend/common"
	"gomall_study/app/frontend/infra/rpc"
	"gomall_study/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {

	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Hot Sale",
		"items": products.Products,
	}, nil

	// resp := make(map[string]any)
	// resp["Title"] = "Hot Sale"
	// item := []map[string]any{
	// 	{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/t-shirt.jpeg"},
	// 	{"Name": "notebook", "Price": 20, "Picture": "/static/image/notebook.jpeg"},
	// 	{"Name": "toiletry-bag", "Price": 180, "Picture": "/static/image/toiletry-bag.jpg"},
	// }
	// resp["Items"] = item
	// return resp, nil
}
