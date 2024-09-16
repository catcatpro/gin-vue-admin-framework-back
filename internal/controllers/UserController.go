package controllers

import (
	"gin_vue_admin_framework/internal/models/requests"
	"gin_vue_admin_framework/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController *UserController) Login(c *gin.Context) {
	var loginInfo requests.LoginRequest
	if c.ShouldBind(&loginInfo) != nil {
		cs := services.CommmonService{}
		_, err := cs.Login(&loginInfo)
		if err != nil {
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"status": "error",
			// 	"msg":    err.Error(),
			// 	"data":   "{}",
			// })
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"msg":    err.Error(),
				"data":   "{}",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"msg":    "success",
			"data":   "{}",
		})
	} else {
		c.JSON(http.StatusInternalServerError, nil)
	}

}
