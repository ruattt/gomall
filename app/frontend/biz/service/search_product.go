package service

import (
	"context"

	product "gomall_study/app/frontend/hertz_gen/frontend/product"
	"gomall_study/app/frontend/infra/rpc"

	rpcproduct "gomall_study/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchProductReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{	
		Query: req.Query,
	})
	if err != nil {
		return nil, err
	}
	
	return utils.H{
		"items": products.Results,
		"query": req.Query,
	}, nil
}
