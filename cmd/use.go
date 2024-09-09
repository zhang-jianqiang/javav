/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhang-jianqiang/javav/pkg"
	"path/filepath"
	"slices"
	"strings"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "切换java版本, 如: javav use jdk-21.0.1",
	Long:  `切换java版本, 如: javav use jdk-21.0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		versionItems, err := pkg.ListVersion()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(args) < 1 {
			fmt.Println("版本为空")
			return
		}

		if !slices.Contains(versionItems, args[0]) {
			fmt.Println("版本号输入错误, 支持的版本如下: \r\n" + strings.Join(versionItems, "\r\n"))
			return
		}

		javaPath, err := pkg.ReadConfig()
		if err != nil {
			fmt.Println("java目录未配置, 请先配置java目录, 例如: javav config D:/programs/Java")
			return
		}

		newJavaEnv := filepath.Join(javaPath, args[0])
		err = pkg.SetLink(newJavaEnv)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Success!\r\nVersion: " + args[0])
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
