package middleware

import (
	Session "gin_vue_admin_framework/pkg/Session/Cookie"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HandleSession(r *gin.Engine) {
	store := Session.InitSession()
	r.Use(sessions.Sessions("systemSession", store))
}
