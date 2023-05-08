package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var cusArgsCheckCmd = &cobra.Command{
	Use: "cusargs",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("至少输入一个参数")
		}
		if len(args) > 2 {
			return errors.New("最多输入两个参数")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		fmt.Println(args)
		fmt.Println("root cmd run end")
	},
}

var argsCheckCmd = &cobra.Command{
	Use:       "args",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"123", "abc", "nick"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root args run begin")
		fmt.Println(args)
		fmt.Println("root args run end")
	},
}

func init() {
	rootCmd.AddCommand(cusArgsCheckCmd)
	rootCmd.AddCommand(argsCheckCmd)
}