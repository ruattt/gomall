package dal

import (
	"gomall_study/app/checkout/biz/dal/mysql"
	"gomall_study/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
