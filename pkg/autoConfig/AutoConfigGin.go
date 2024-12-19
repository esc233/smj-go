package autoConfig

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"os"
)

type serverConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

func InitRouter(config serverConfig) {
	log.Info().Msg("Starting the application...")

	if err := os.Setenv("GIN_MODE", config.Mode); err != nil {
		panic(err)
	}

	r := gin.Default()
	if err := r.Run(":" + config.Port); err != nil {
		panic(err) // 启动失败时退出程序并输出错误
	}
	log.Info().Msg("Application started successfully!")
}
