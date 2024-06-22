package global

import (
	"database_lesson/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config      *config.Config
	DB          *gorm.DB
	Log         *logrus.Logger
	RedisClient *redis.Client
)
