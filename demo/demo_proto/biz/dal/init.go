package dal

import (
	"demo/demo_proto/biz/dal/mysql"
	// "demo/demo_proto/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
