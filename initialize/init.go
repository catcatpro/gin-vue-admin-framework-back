package initialize

import (
	"gin_vue_admin_framework/configs"
)

// 初始化系统
func InitSystem() {
	{
		configs.InitConfig()
		initDB()
		initRedis()
	}
}
