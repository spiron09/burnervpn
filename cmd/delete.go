/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spiron09/burnervpn/internal/client"
)

// disconnectCmd represents the disconnect command
var disconnectCmd = &cobra.Command{
	Use:   "delete <session-id>",
	Short: "Deletes a specified VPN server",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sessionID := args[0]
		sessions, err := loadMetadata()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		session, ok := sessions[sessionID]
		if !ok {
			fmt.Fprintf(os.Stderr, "Error: Session %s not found\n", sessionID)
			os.Exit(1)
		}

		configFilePath := filepath.Join(os.Getenv("HOME"), ".burnervpn", "sessions", session.FileName)
		client := client.NewClient()
		resp, err := client.DeleteSession(sessionID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		err = os.Remove(configFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Delete the session file
		delete(sessions, sessionID)
		err = updateMetadata(sessions)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Session %s deleted successfully\n", sessionID)
		fmt.Printf("\n-----USAGE-----\n")
		fmt.Printf("Time: %v\n", resp.DurationInSeconds)
		fmt.Printf("Cost: %v\n", resp.Cost)
	},
}

func init() {
	rootCmd.AddCommand(disconnectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disconnectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disconnectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
