/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	svcs "nexus-sds.com/deploydock/services"
	u "nexus-sds.com/deploydock/utilities"

	"github.com/spf13/cobra"
)

// bootstrapCmd represents the bootstrap command
var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Setup new deploydock instance",
	Long:  `Setup a new deploy dock instance`,
	Run: func(cmd *cobra.Command, args []string) {
		ddkConfig := u.LoadConfig()
		url := u.CreateUrl(ddkConfig)
		svcs.SetupDeployDockGitRepo(ddkConfig, url)
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bootstrapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bootstrapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
