package common

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	COM_DB    *gorm.DB
	COM_REDIS *redis.Client
)
