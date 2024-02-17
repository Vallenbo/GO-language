package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 把当前目录加入到配置文件的搜索路径中
	viper.SetConfigName("\\viper_use\\config.yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())

	// 读取配置值
	fmt.Printf("username: %s\n", viper.Get("name"))
	fmt.Printf("username: %s\n", viper.Get("mysql"))
}
