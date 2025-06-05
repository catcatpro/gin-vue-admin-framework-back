package middlewares

import (
	"gin_vue_admin_framework/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录鉴权中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		//验证token是否有效？
		_, err := j.ParseToken(token, false)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"msg":    err.Error(),
				"data":   "{}",
			})
			return
		}
		//验证token是否过期？
		_, err = j.VerifyTokenExpiresAt(token, false)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"msg":    err.Error(),
				"data":   "{''}",
			})
			return
		}
		c.Next()
	}
}
