/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hiCmd represents the hi command
var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "Debug command.",
	Long:  `Debug command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from rdvl-cli")
	},
}

func init() {
	rootCmd.AddCommand(hiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
