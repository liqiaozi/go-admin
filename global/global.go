package global

import (
	"github.com/spf13/viper"
	"lixuefei.com/go-admin/config"
)

type globalVars struct {
	ConfigViper *viper.Viper
	Application config.Application
}

var App = new(globalVars)
