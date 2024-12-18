package config

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitZeroLog() {
	logWriter := &lumberjack.Logger{
		Filename:   C.Log.Filename,   // 日志文件名
		MaxSize:    C.Log.MaxSize,    // 每个文件最大大小（MB）
		MaxBackups: C.Log.MaxBackups, // 保留旧日志文件的最大数量
		MaxAge:     C.Log.MaxAge,     // 文件最大保存天数
		Compress:   C.Log.Compress,   // 是否压缩旧日志文件
	}
	// 设置 zerolog 的输出为 lumberjack
	log.Logger = zerolog.New(logWriter).With().Timestamp().Logger()
}
