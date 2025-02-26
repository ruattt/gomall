package rpc

import (
	"sync"
	// "sync"

	"gomall_study/app/frontend/conf"
	frontendUtils "gomall_study/app/frontend/utils"
	"gomall_study/rpc_gen/kitex_gen/cart/cartservice"
	"gomall_study/rpc_gen/kitex_gen/checkout/checkoutservice"
	"gomall_study/rpc_gen/kitex_gen/order/orderservice"
	"gomall_study/rpc_gen/kitex_gen/product/productcatalogservice"
	"gomall_study/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	once       sync.Once
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	CartClient cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient orderservice.Client
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}


func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}


func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	// cbs.UpdateServiceCBConfig("shop-frontend/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

	// opts = append(opts, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
	// 	methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
	// 	if err == nil {
	// 		return resp, err
	// 	}
	// 	if methodName != "ListProducts" {
	// 		return resp, err
	// 	}
	// 	return &product.ListProductsResp{
	// 		Products: []*product.Product{
	// 			{
	// 				Price:       6.6,
	// 				Id:          3,
	// 				Picture:     "/static/image/t-shirt.jpeg",
	// 				Name:        "T-Shirt",
	// 				Description: "CloudWeGo T-Shirt",
	// 			},
	// 		},
	// 	}, nil
	// }))))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("Checkout", opts...)
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleError(err)
}
