package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short desc",
	Long:  "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.PersistentFlags().Lookup("author").Value,
			cmd.PersistentFlags().Lookup("config").Value,

			cmd.PersistentFlags().Lookup("License").Value,
			cmd.Flags().Lookup("source").Value,
		)
		fmt.Println(
			viper.GetString("author"),
			viper.GetString("license"),
		)
		fmt.Println("root cmd run end")
	},
	TraverseChildren: true,
}

func Execute() {
	rootCmd.Execute()
}

var cfgFile, userLicense string

func init() {
	cobra.OnInitialize(initConfig) //设置在调用每个命令的 Execute 方法时要运行的传递函数
	//rootCmd.PersistentFlags() //持久化标志，可以传递子命令
	//rootCmd.Flags()           //标志，只有当前命令可以使用
	rootCmd.Flags().StringP("source", "s", "", "") //添加本地标识

	rootCmd.PersistentFlags().Bool("viper", true, "")                          //按名称接收命令行参数
	rootCmd.PersistentFlags().StringP("author", "a", "Your Name", "")          //指定flag缩写
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")            //通过指针，将值赋值到字段
	rootCmd.PersistentFlags().StringVarP(&userLicense, "License", "l", "", "") //通过指针，将值赋值到字段,并指定flag缩写-l

	//rootCmd.PersistentFlags().BoolP("viper", "v", true, "")
	//var vbool bool
	//rootCmd.PersistentFlags().BoolVarP(&vbool, "viper", "v", true, "") //BoolVarP 类似于 BoolVar，但接受一个可以在短划线后使用的速记字母。
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("license", rootCmd.PersistentFlags().Lookup("license"))
	viper.SetDefault("author", "default author") //设置默认值
	viper.SetDefault("license", "default license")
}

func initConfig() { //配置文件初始化,命令值从配置文件中获取
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserConfigDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv() //检查环境变量，将配置的键值加载到viper
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Using config file : ", viper.ConfigFileUsed())
}