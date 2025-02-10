package service

import (
	"context"

	common "gomall_study/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	resp := make(map[string]any)

	resp["Title"] = "Hot Sale"

	item := []map[string]any{
		{"Name": "T-shirt", "Price": 100, "Picture": "/static/image/t-shirt.jpeg"},
		{"Name": "notebook", "Price": 20, "Picture": "/static/image/notebook.jpeg"},
		{"Name": "toiletry-bag", "Price": 180, "Picture": "/static/image/toiletry-bag.jpg"},
	}
	resp["Items"] = item

	return resp, nil
}
