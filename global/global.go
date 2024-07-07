package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/config"
)

type globalVars struct {
	ConfigViper *viper.Viper
	Server      config.Server
	DB          *gorm.DB
	RedisClient redis.UniversalClient
}

var App = new(globalVars)
