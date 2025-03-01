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
	"gomall_study/common/clientsuite"
)

var (
	once           sync.Once
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	ServiceName    = frontendUtils.ServiceName
	MetricsPort    = conf.GetConf().Hertz.MetricsPort
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
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
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("Checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendUtils.MustHandleError(err)
}
