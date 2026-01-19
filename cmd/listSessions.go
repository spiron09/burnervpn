/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listSessionsCmd represents the listSessions command
var listSessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "List saved session configurations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sessions, err := loadMetadata()
		if err != nil {
			fmt.Printf("Error loading sessions: %v\n", err)
			return
		}

		if len(sessions) == 0 {
			fmt.Println("No sessions found.")
			return
		}

		fmt.Println("Saved Sessions:")
		for _, session := range sessions {
			fmt.Printf("ID: %s, Region: %s, Created: %s\n", session.SessionID, session.Region, session.CreatedAt)
		}
	},
}

func init() {
	// rootCmd.AddCommand(listSessionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listSessionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listSessionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
