/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	svcs "nexus-sds.com/deploydock/services"
	u "nexus-sds.com/deploydock/utilities"

	"github.com/spf13/cobra"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh your deploydock server instance",
	Long:  `Refresh the git repository that your deploy dock instance uses`,
	Run: func(cmd *cobra.Command, args []string) {
		ddkConfig := u.LoadConfig()
		svcs.UpdateDeployDockGitRepo(ddkConfig)
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refreshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// refreshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
