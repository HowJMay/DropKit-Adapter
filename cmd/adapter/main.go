package main

import (
	"fmt"

	"github.com/DropKit/DropKit-Adapter/configs"
	"github.com/DropKit/DropKit-Adapter/internal/routes"
	"github.com/DropKit/DropKit-Adapter/pkg/logging"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, err := logging.NewLogger(logging.DEBUG, "main")
	if err != nil {
		logger.Debug("fail to create a logger", zap.Error(err))
	}

	viper.SetConfigName("config")
	viper.AddConfigPath("../../configs")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		logger.Debug("can't read config file", zap.Error(err))
	}
	if err := viper.Unmarshal(&configs.Configs); err != nil {
		logger.Debug("can't unmarshal config file", zap.Error(err))
	}

	// service, err := services.DependencyServicesCheck()
	// if err != nil {
	// 	logger.FatalDependencyService(service, err)
	// 	return
	// }

	// logger.InfoDependencyService()

	router := routes.SetupRouter()
	fmt.Println(":" + viper.GetString(`DROPKIT.PORT`))
	router.Run(":" + viper.GetString(`DROPKIT.PORT`))
	// router.Run(":8080")
	fmt.Println(logger)
}
