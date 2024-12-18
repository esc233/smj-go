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
			V.SetConfigFile(filepath.Join(configDir, file.Name()))
			err := V.MergeInConfig()
			if err != nil {
				fmt.Printf("加载配置文件失败: %s, 错误: %v\n", file.Name(), err)
				return
			} else {
				fmt.Printf("成功加载配置文件: %s\n", file.Name())
			}
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件夹失败: %v\n", err)
		return
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("读取配置文件夹失败: %v\n", err)
		return
	} else {
		isLoad = true
	}
}
