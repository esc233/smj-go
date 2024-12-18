package config

type appInfo struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Config struct {
	Server  serverConfig `mapstructure:"server"`
	Log     logConfig    `mapstructure:"log"`
	AppInfo appInfo      `mapstructure:"app_info"`
}

type logConfig struct {
	Filename   string `mapstructure:"filename"`     // 日志文件名
	MaxSize    int    `mapstructure:"max_size"`     // 每个文件最大大小（MB）
	MaxBackups int    `mapstructure:"max_back_ups"` // 保留旧日志文件的最大数量
	MaxAge     int    `mapstructure:"max_age"`      // 文件最大保存天数
	Compress   bool   `mapstructure:"compress"`     // 是否压缩旧日志文件
}

type serverConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
