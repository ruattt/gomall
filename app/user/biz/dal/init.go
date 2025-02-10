package dal

import (
	"gomall_study/app/user/biz/dal/mysql"
	"gomall_study/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
