/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	u "nexus-sds.com/deploydock/utilities"
	"os"
	"path"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [compose-app]",
	Short: "Create new deploydock server compose app",
	Long:  `Use this command to create a new compose app`,
	Run: func(cmd *cobra.Command, args []string) {
		ddkConfig := u.LoadConfig()

		if len(args) == 0 {
			// Exit out of loop
			fmt.Println("Name of compose app is required")
			os.Exit(1)
		}
		if u.FileExists(path.Join(ddkConfig.ServerRoot, args[0])) {
			fmt.Println("Compose app already exists")
			os.Exit(1)
		}
		err := os.MkdirAll(path.Join(ddkConfig.ServerRoot, args[0]), 0644)
		if err != nil {
			fmt.Printf("Error creating compose app: %s\n", err)
			return
		}

		// Create Compose.yaml
		_, err = os.Create(path.Join(ddkConfig.ServerRoot, args[0], "compose.yaml"))

		fmt.Println("Compose app created")

	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
