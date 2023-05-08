package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "add",
	Short: "short init",
	Long:  "Long init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		fmt.Println(
			cmd.Flags().Lookup("viper").Value,
			cmd.Flags().Lookup("author").Value,
			cmd.Flags().Lookup("config").Value,
			viper.GetString("author"),
			cmd.Flags().Lookup("License").Value,
			cmd.Parent().Flags().Lookup("source").Value,
		)
		fmt.Println("root cmd run end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}