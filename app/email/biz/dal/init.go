package dal

import (
	"gomall_study/app/email/biz/dal/mysql"
	"gomall_study/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
