package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"gomall_study/app/cart/conf"
	cartutils "gomall_study/app/cart/utils"
	"gomall_study/common/clientsuite"
	"gomall_study/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	// ServiceName   = "cart"
	// RegistryAddr  = "127.0.0.1:8500"
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
