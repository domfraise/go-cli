/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// newbranchCmd represents the newbranch command
var newbranchCmd = &cobra.Command{
	Use:   "newbranch",
	Short: "Create and checkout a new branch",
	Long:  `Stash any current changes, checks out either main or master, pulls, then creates new branch`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.PlainOpen("./") //this
		if err != nil {
			log.Fatal(err)
		}
		hasMain := hasMainBranch(repo)
		fmt.Println(hasMain)

	},
}

func hasMainBranch(repo *git.Repository) (hasMain bool) {
	branches, _ := repo.Branches()
	branches.ForEach(func(c *plumbing.Reference) error {
		fmt.Println(c)
		if strings.Contains(c.String(), "refs/heads/main") {
			hasMain = true
		}
		return nil
	})
	return
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
