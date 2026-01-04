/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spiron09/burnervpn/internal/client"
)

type SessionMetadata struct {
	SessionID string `json:"session_id"`
	FileName  string `json:"file_name"`
	Region    string `json:"region"`
	CreatedAt string `json:"created_at"`
}

func loadMetadata() ([]SessionMetadata, error) {
	path := filepath.Join(os.Getenv("HOME"), ".burnervpn", "metadata.json")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return make([]SessionMetadata, 0), nil
	}
	if err != nil {
		return nil, err
	}
	var sessions []SessionMetadata
	err = json.Unmarshal(data, &sessions)
	return sessions, err
}

func saveMetadata(session SessionMetadata) error {
	path := filepath.Join(os.Getenv("HOME"), ".burnervpn", "metadata.json")
	sessions, err := loadMetadata()
	if err != nil {
		return err
	}
	sessions = append(sessions, session)
	data, err := json.MarshalIndent(sessions, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func saveConfig(SessionID, region, config string) (string, error) {
	dir := filepath.Join(os.Getenv("HOME"), ".burnervpn", "sessions")

	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s_%s.conf", SessionID[:8], region)
	err = saveMetadata(SessionMetadata{
		SessionID: SessionID,
		FileName:  filename,
		Region:    region,
		CreatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		return "", err
	}
	configFilePath := filepath.Join(dir, filename)
	return configFilePath, os.WriteFile(configFilePath, []byte(config), 0600)
}

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "create <region>",
	Short: "Create a VPN server in the specified region",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		region := args[0]
		fmt.Printf("Creating a VPN server in %s...\n", region)

		c := client.NewClient()
		resp, err := c.CreateSession(region)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		configFilePath, err := saveConfig(resp.SessionID, region, resp.WireGuardConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Created a VPN server in %s\n", region)
		fmt.Printf("Config file saved to %s\n", configFilePath)
		fmt.Printf("Metadata saved to %s\n", filepath.Join(os.Getenv("HOME"), ".burnervpn", "metadata.json"))
		//TODO: print QR Code
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
