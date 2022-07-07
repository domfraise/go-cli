/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newbranchCmd represents the newbranch command
var newbranchCmd = &cobra.Command{
	Use:   "newbranch",
	Short: "Create and checkout a new branch",
	Long:  `Stash any current changes, checks out either main or master, pulls, then creates new branch`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("newbranch called")
	},
}

func init() {
	rootCmd.AddCommand(newbranchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newbranchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newbranchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
