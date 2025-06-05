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
	token, refresh_token, err := cs.Login(&loginInfo)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    err.Error(),
			"data":   "{}",
		})
		return
	}
	//c.Writer.Header().Set("authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "success",
		"data":   gin.H{"token": token, "refresh_token": refresh_token},
	})
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	refreshToken := requests.RefreshTokenRequest{}
	err := c.ShouldBind(&refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(402, gin.H{
			"status": "error",
			"msg":    "Parameter error",
			"data":   "{}",
		})

		return
	}
	cs := services.SysUserService{}

	token, err := cs.RefreshToken(&refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(402, gin.H{
			"status": "error",
			"msg":    err.Error(),
			"data":   "{}",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "success",
		"data":   gin.H{"token": token},
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
	token := c.GetHeader("x-header-token")
	data.Data = token
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
