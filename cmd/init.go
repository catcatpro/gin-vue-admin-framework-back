package cmd

import (
	"gin_vue_admin_framework/configs"
	"gin_vue_admin_framework/utils"
)

// 初始化系统
func InitSystem() {
	{
		configs.InitConfig()
		utils.InitDB()
		utils.InitRedis()
		utils.InitRedis()
	}
}
