package dal

import (
	"gomall_study/demo/front/biz/dal/mysql"
	"gomall_study/demo/front/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
