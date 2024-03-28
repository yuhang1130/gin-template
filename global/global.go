package global

import (
	"gin-template/config"

	"github.com/boj/redistore"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config     *config.AllConfig // 全局Config
	Log        *log.Logger
	DB         *gorm.DB
	Redis      *redis.Client
	RedisStore *redistore.RediStore
)
