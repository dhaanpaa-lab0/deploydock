/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// submitCmd represents the submit command
var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit changes to your deploydock server instances to git repository",
	Long:  `Submit changes to your deploydock server instances to git repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		message := cmd.Flag("message").Value.String()
		if message == "" {
			
		}
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	submitCmd.Flags().StringP("message", "m", "", "Commit message")
}
