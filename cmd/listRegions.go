/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiron09/burnervpn/internal/client"
)

// listRegionsCmd represents the listRegions command
var listRegionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "Lists the available server locations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		regions, err := c.ListRegions()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Available regions:")
		for _, region := range regions.Regions {
			fmt.Printf("- %s (%s)\n", region.Name, region.Slug)
		}
	},
}

func init() {
	// rootCmd.AddCommand(listRegionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listRegionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listRegionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
