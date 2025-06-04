package adminControllers

import (
	"fmt"
	"gin_vue_admin_framework/internal/models"
	"gin_vue_admin_framework/internal/models/requests"
	"gin_vue_admin_framework/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) LoginAction(c *gin.Context) {
	loginInfo := requests.LoginRequest{}
	err := c.ShouldBind(&loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Parameter error",
			"data":   "{}",
		})

		return
	}
	fmt.Println(loginInfo)
	cs := services.SysUserService{}
	token, err := cs.Login(&loginInfo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    err.Error(),
			"data":   "{}",
		})
		return
	}
	c.Writer.Header().Set("authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "success",
		"data":   "{}",
	})

}

func (uc *UserController) CreateUserAction(c *gin.Context) {
	createUser := requests.CreateUserRequest{}
	err := c.ShouldBind(&createUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Parameter error",
			"data":   "{}",
		})

		return
	}

	cs := services.SysUserService{}
	err = cs.CreateUser(&createUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    err.Error(),
			"data":   "{}",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "success",
		"data":   "{}",
	})

}

func (uc *UserController) GetUserInfoAction(c *gin.Context) {
	var err error
	data := requests.GetUserInfoRequest{}
	err = c.ShouldBind(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Parameter error",
			"data":   "{}",
		})

		return
	}

	var userInfo models.User
	cs := services.SysUserService{}
	userInfo, err = cs.GetUserInfo(&data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    "Parameter error",
			"data":   "{}",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "success",
		"data":   userInfo,
	})

}
