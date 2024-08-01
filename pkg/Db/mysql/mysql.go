package mysql

import (
	"fmt"
	"gin_vue_admin_framework/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	//处理配置

	config := configs.SystemConfigs.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
	fmt.Println(dsn)

	//连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connect failed, check your database configuration. Error: " + err.Error())
	}
	return db
}

func init() {
	Db = Connect()
}
