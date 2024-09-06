/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhang-jianqiang/javav/pkg"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "多个java安装包存放目录,示例: javav config D:/programs/Java",
	Long:  `多个java安装包存放目录`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("java安装包存放目录不能为空")
			return
		}

		err := pkg.SetConfig(args[0])
		if err != nil {
			fmt.Println("设置配置失败,请检查是否有权限")
			return
		}

		fmt.Println("Config Success!\r\nPATH: " + args[0])
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
