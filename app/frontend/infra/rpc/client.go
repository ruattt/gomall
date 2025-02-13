package rpc

import (
	"sync"
	// "sync"

	"gomall_study/app/frontend/conf"
	frontendUtils "gomall_study/app/frontend/utils"
	"gomall_study/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	once       sync.Once
	UserClient userservice.Client
)

func InitClient() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
