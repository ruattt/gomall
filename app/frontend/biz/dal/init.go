package dal

import (
	"app/frontend/biz/dal/mysql"
	"app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
