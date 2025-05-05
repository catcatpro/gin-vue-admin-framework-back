package initialize

import (
	"fmt"
	"gin_vue_admin_framework/common"
	"gin_vue_admin_framework/configs"
	"gin_vue_admin_framework/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DbInterface interface {
	connect() *gorm.DB
}

type MysqlDB struct{}

func (m *MysqlDB) connect() *gorm.DB {
	//处理配置
	config := configs.SystemConfigs.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database)

	//连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Prefix, // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,          // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic("Connect failed, check your database configuration. Error: " + err.Error())
	}

	db.AutoMigrate(&models.User{})
	return db
}

func connectIng() *gorm.DB {
	var db DbInterface = new(MysqlDB)
	return db.connect()

}

func initDB() {
	common.COM_DB = connectIng()
	// fmt.Println("common.COM_DB", common.COM_DB)
}
