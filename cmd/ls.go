/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zhang-jianqiang/javav/pkg"
	"strings"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "java版本列表, 示例: javav ls",
	Long:  `java版本列表, 示例: javav ls`,
	Run: func(cmd *cobra.Command, args []string) {
		javaversionItems, err := pkg.ListVersion()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("java版本: \r\n" + strings.Join(javaversionItems, "\r\n"))
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
