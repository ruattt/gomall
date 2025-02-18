package dal

import (
	"gomall_study/app/product/biz/dal/mysql"
	// "gomall_study/app/product/biz/dal/redis"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
