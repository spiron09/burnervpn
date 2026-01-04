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

// usageCmd represents the usage command
var usageCmd = &cobra.Command{
	Use:   "usage <session-id>",
	Short: "Shows the usage of a specified session",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sessionID := args[0]

		sessions, err := loadMetadata()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		_, ok := sessions[sessionID]
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: Session %s not found\n", sessionID)
			os.Exit(1)
		}

		c := client.NewClient()
		resp, err := c.GetUsage(sessionID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Print usage
		fmt.Printf("\n-----USAGE-----\n")
		fmt.Printf("Session: %s\n", sessionID)
		fmt.Printf("Duration: %.2f seconds\n", resp.DurationInSeconds)
		fmt.Printf("Cost: $%.2f\n", resp.Cost)

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
