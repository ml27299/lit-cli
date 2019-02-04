package cmd

import (
	//"fmt"
	"gopkg.in/src-d/go-git.v4"
	"github.com/spf13/cobra"
	. "../helpers"
	"../helpers/paths"
	//"errors"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds files to git repo, analogous to \"git add\"",
	Long: `ex. lit add .`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {

	dir, err := paths.FindRootDir()
	CheckIfError(err)
	
	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	links, err := FindHardLinkedFilePaths()
	CheckIfError(err)

	err = UpdateGitignore(dir, links)
	CheckIfError(err)

	for i := 0; i < len(submodules); i++ {

		worktree, err := GetSubmoduleWorkTree(submodules[i])
		CheckIfError(err)

		worktree.Add(args[0])
	}

	repo, err := git.PlainOpen(dir)
	CheckIfError(err)

	worktree, err := repo.Worktree()
	CheckIfError(err)

	worktree.Add(args[0])
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("foo", "", "A help for foo")
	addCmd.Flags().BoolP("toggle", "t", false, "works exactly like git add, but does it for every repo")
}
