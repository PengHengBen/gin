package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

var V *viper.Viper

func Init() {
	// 定义配置文件路径
	config := "./config/config.toml"
	// 初始化 viper
	V = viper.New()
	V.SetConfigFile(config)
	V.SetConfigType("toml")
	// 处理读取文件错误
	if err := V.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}
}
