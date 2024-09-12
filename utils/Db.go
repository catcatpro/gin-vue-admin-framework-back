package utils

import (
	"fmt"
	"gin_vue_admin_framework/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbInterface interface {
	connect() *gorm.DB
}

var Db *gorm.DB

type MysqlDB struct{}

func (m MysqlDB) connect() *gorm.DB {
	//处理配置
	config := configs.SystemConfigs.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
	// fmt.Println(dsn)

	//连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Connect failed, check your database configuration. Error: " + err.Error())
	}
	return db
}

func connectIng() *gorm.DB {
	var db DbInterface

	db = MysqlDB{}

	return db.connect()
}

func InitDB() {
	Db = connectIng()
}
