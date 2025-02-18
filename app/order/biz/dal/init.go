package dal

import (
	"gomall_study/app/order/biz/dal/mysql"
	"gomall_study/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
