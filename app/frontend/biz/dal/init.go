package dal

import (
	"gomall_study/app/frontend/biz/dal/mysql"
	"gomall_study/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
