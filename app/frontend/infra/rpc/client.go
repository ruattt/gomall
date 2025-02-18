package rpc

import (
	"sync"
	// "sync"

	"gomall_study/app/frontend/conf"
	frontendUtils "gomall_study/app/frontend/utils"
	"gomall_study/rpc_gen/kitex_gen/cart/cartservice"
	"gomall_study/rpc_gen/kitex_gen/checkout/checkoutservice"
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
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
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