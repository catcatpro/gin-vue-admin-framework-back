package middlewares

import "github.com/gin-gonic/gin"

// 登录鉴权中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO 登录鉴权中间件
	}
}
