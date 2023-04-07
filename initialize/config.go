package initialize

import (
	"admin_api/global"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Pre        = "./etc/config_pre.yaml"
	Prod       = "./etc/config_prod.yaml"
	configFile = Prod
)

func Config() {
	viper.AutomaticEnv()
	if debug := viper.GetBool("ADMIN_API_DEBUG"); debug {
		configFile = Pre
	}

	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("server config info: %v", global.ServerConfig)
}
