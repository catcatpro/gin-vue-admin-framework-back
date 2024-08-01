package Session

import (
	"gin_vue_admin_framework/configs"
	"github.com/gin-contrib/sessions/cookie"
)

func InitSession() cookie.Store {
	sessionConfig := configs.SystemConfigs.Session
	store := cookie.NewStore([]byte(sessionConfig.Secret))
	return store
}
