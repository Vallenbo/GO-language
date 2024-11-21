package viper_use

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func Test_viper(*testing.T) {
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
	fmt.Printf("key: %s\n", viper.Get("key.0")) // 读取一个包含多个值的键
}

// 动态监听文件
func Test_viperListenConf(t *testing.T) {
	viper.SetConfigFile(`./config.yaml`)
	viper.ReadInConfig()

	// 注册每次配置文件发生变更后都会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed: %s\n", e.Name)
	})

	// 监控并重新读取配置文件，需要确保在调用前添加了所有的配置路径
	viper.WatchConfig()

	// 阻塞程序，这个过程中可以手动去修改配置文件内容，观察程序输出变化
	time.Sleep(time.Second * 20)

	// 读取配置值
	fmt.Printf("username: %s\n", viper.Get("key"))
}
