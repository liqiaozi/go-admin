package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"lixuefei.com/go-admin/config"
)

type globalVars struct {
	ConfigViper *viper.Viper
	Application config.Application
	DB          *gorm.DB
}

var App = new(globalVars)
