/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// usageCmd represents the usage command
var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Shows the current session's usage of your VPN connection",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("usage called")
	},
}

func init() {
	rootCmd.AddCommand(usageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
