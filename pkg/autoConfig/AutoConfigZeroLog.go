package autoConfig

import (
	"encoding/json"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type logConfig struct {
	Filename   string `mapstructure:"filename" json:"filename"`       // 日志文件名
	MaxSize    int    `mapstructure:"max_size" json:"maxSize"`        // 每个文件最大大小（MB）
	MaxBackups int    `mapstructure:"max_back_ups" json:"maxBackups"` // 保留旧日志文件的最大数量
	MaxAge     int    `mapstructure:"max_age" json:"maxAge"`          // 文件最大保存天数
	Compress   bool   `mapstructure:"compress" json:"compress"`       // 是否压缩旧日志文件
	Console    bool   `mapstructure:"console" json:"console"`         // 是否输出到控制台
}

func (c *logConfig) fillLogConfig() {
	if c.Filename == "" {
		c.Filename = "logs/app.log"
	}
	if c.MaxSize == 0 {
		c.MaxSize = 100
	}
	if c.MaxBackups == 0 {
		c.MaxBackups = 50
	}
	if c.MaxAge == 0 {
		c.MaxAge = 7
	}
}

func InitZeroLog(logConfig logConfig) {
	logConfig.fillLogConfig()
	logWriter := &lumberjack.Logger{
		Filename:   logConfig.Filename,   // 日志文件名
		MaxSize:    logConfig.MaxSize,    // 每个文件最大大小（MB）
		MaxBackups: logConfig.MaxBackups, // 保留旧日志文件的最大数量
		MaxAge:     logConfig.MaxAge,     // 文件最大保存天数
		Compress:   logConfig.Compress,   // 是否压缩旧日志文件
	}
	userJSON, _ := json.Marshal(logConfig)
	log.Logger = zerolog.New(logWriter).With().Timestamp().Logger()
	if logConfig.Console {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	log.Info().Msgf("log 初始化成功 %s", string(userJSON))
}
