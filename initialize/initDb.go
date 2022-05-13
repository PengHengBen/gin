package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDb() {
	// 获取数据库连接配置信息
	host := V.Get("mysql.host")
	port := V.Get("mysql.port")
	user := V.Get("mysql.user")
	password := V.Get("mysql.password")
	database := V.Get("mysql.database")
    // 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",user,password,host,port,database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 处理数据库连接错误
	if err != nil {
		panic(err)
	}
}
