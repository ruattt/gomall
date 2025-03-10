package mysql

import (
	"fmt"
	"os"

	"gomall/app/user/biz/model"
	"gomall/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// conf.GetConf().MySQL.DSN
	// "%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err!= nil{
		panic(err)
	}
	DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
