package dal

import (
	"gomall_study/app/payment/biz/dal/mysql"
	"gomall_study/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
