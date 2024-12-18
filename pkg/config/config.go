package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var C Config
var V = viper.New()
var isLoad = false

func IsLoad() bool {
	return isLoad
}
func hasValidExtension(filename string) bool {
	// 检查文件是否以指定后缀结尾
	return strings.HasSuffix(strings.ToLower(filename), ".yaml") ||
		strings.HasSuffix(strings.ToLower(filename), ".yml") ||
		strings.HasSuffix(strings.ToLower(filename), ".json")
}

func LoadConfig(configDir string) {
	if strings.TrimSpace(configDir) == "" {
		configDir = "./config"
	}
	files, err := os.ReadDir(configDir)
	if err != nil {
		fmt.Printf("读取配置文件夹失败: %v\n", err)
		return
	}

	for _, file := range files {
		// 忽略子目录
		if file.IsDir() {
			continue
		}

		// 获取文件扩展名
		if hasValidExtension(file.Name()) {
			// 打印文件名
			fmt.Printf("找到配置文件: %s\n", file.Name())
			filename := file.Name()
			baseName := filename[:len(filename)-len(filepath.Ext(filename))] // 去掉扩展名
			// 创建一个临时 Viper 实例加载当前文件
			tempViper := viper.New()
			tempViper.SetConfigName(baseName)
			tempViper.AddConfigPath(configDir)
			tempViper.SetConfigType(filepath.Ext(filename)[1:]) // 动态设置配置类型
			// 读取当前配置文件
			if err := tempViper.ReadInConfig(); err != nil {
				fmt.Printf("加载配置文件失败 (%s): %v", filename, err)
				return
			}
			// 使用 MergeInConfig 将文件内容合并到主 Viper 实例
			if err := V.MergeConfigMap(tempViper.AllSettings()); err != nil {
				fmt.Printf("合并配置文件失败 (%s): %v", filename, err)
				return
			}
		}
	}
	err = V.Unmarshal(&C)
	if err != nil {
		fmt.Printf("读取配置文件夹失败: %v\n", err)
		return
	} else {
		isLoad = true
	}
}
