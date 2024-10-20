package middlewares

import (
	"gin_vue_admin_framework/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录鉴权中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO 登录鉴权中间件
		token := c.GetHeader("x-header-token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"msg":    "token not null.",
				"data":   "{}",
			})
			return
		}
		j := utils.NewJWT()
		_, err := j.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"msg":    err.Error(),
				"data":   "{}",
			})
			return
		}
		c.Next()
	}
}
