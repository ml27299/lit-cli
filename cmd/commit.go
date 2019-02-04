package cmd

import (
	"time"
	"gopkg.in/src-d/go-git.v4"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	. "../helpers"
	"../helpers/paths"
)

var (
	message string
	all bool

	commitCmd = &cobra.Command{
		Use:   "commit",
		Short: "commit files to git repo, analogous to \"git commit\"",
		Long: `ex. lit commit -am "my commit message"`,
		Run: commitRun,
	}
)

func commitRun(cmd *cobra.Command, args []string) {

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	links, err := FindHardLinkedFilePaths()
	CheckIfError(err)

	err = UpdateGitignore(dir, links)
	CheckIfError(err)

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)
		Info("Entering "+*&status.Path+"...")

		worktree, err := GetSubmoduleWorkTree(submodules[i])
		CheckIfError(err)

		
		_, err = worktree.Commit(message, &git.CommitOptions{
			All: all,
			Author: &object.Signature{
				When: time.Now(),
			},
		})
		CheckIfError(err)
	}

	Info("Entering /...")
	repo, err := git.PlainOpen(dir)
	CheckIfError(err)

	worktree, err := repo.Worktree()
	CheckIfError(err)

	_, err = worktree.Commit(message, &git.CommitOptions{
		All: all,
		Author: &object.Signature{
			When: time.Now(),
		},
	})
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.PersistentFlags().String("foo", "", "A help for foo")
	commitCmd.Flags().BoolVarP(&all, "all", "a", false, "Tell the command to automatically stage files that have been modified and deleted, but new files you have not told Git about are not affected.")
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "The message for the commit")
}
