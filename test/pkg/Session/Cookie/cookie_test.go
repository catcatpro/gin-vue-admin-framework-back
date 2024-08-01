package Session

import (
	"gin_vue_admin_framework/configs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func initSession(r *gin.Engine) {
	sessionConfig := configs.SystemConfigs.Session
	store := cookie.NewStore([]byte(sessionConfig.Secret))
	r.Use(sessions.Sessions("systemSession", store))
}
