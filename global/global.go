package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/bootstrap/config"
)

type globalVars struct {
	ConfigViper *viper.Viper          `json:"config_viper,omitempty"`
	Server      config.Server         `json:"server"`
	DB          *gorm.DB              `json:"db,omitempty"`
	RedisClient redis.UniversalClient `json:"redis_client,omitempty"`
	BlackCache  local_cache.Cache     `json:"black_cache,omitempty"`
}

var App = new(globalVars)
