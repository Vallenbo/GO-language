package viper_use

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

type (
	conf struct {
		Name      string    `mapstructure:"name"`
		Key       []string  `mapstructure:"key"`
		MysqlConf MysqlConf `mapstructure:"mysql"`
	}
	MysqlConf struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	}
)

func Test_confReflect(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("./config.yaml") // 指定配置文件路径
	err := v.ReadInConfig()          // 读取配置信息
	if err != nil {                  // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	Conf := conf{}
	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	fmt.Printf("conf :%#v \n", Conf)
}
