package cmd

import (
	"gopkg.in/src-d/go-git.v4"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
	. "../helpers"
	"../helpers/paths"
)

var (
	branch_name string
	checkoutCmd = &cobra.Command{
		Use:   "checkout",
		Short: "checkout branch from git repo, analogous to \"git checkout\"",
		Long: `ex. lit checkout master`,
		Run: checkoutRun,
	}
)

func checkoutRun(cmd *cobra.Command, args []string) {

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	var (
		branch_ref plumbing.ReferenceName
		create = false
		name string
	)

	if branch_name != "" {

		name = branch_name
		create = true

	} else {
		name = args[0]
	}

	branch_ref = plumbing.NewBranchReferenceName(name)	
	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)
		Info("Entering "+*&status.Path+"...")

		worktree, err := GetSubmoduleWorkTree(submodules[i])
		CheckIfError(err)

		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: branch_ref,
			Create: create,
		})

		if err != nil && err == plumbing.ErrReferenceNotFound {
			err = worktree.Checkout(&git.CheckoutOptions{
				Branch: branch_ref,
				Create: true,
			})
		}

		CheckIfError(err)
	}

	Info("Entering /...")
	repo, err := git.PlainOpen(dir)
	CheckIfError(err)

	worktree, err := repo.Worktree()
	CheckIfError(err)

	err = worktree.Checkout(&git.CheckoutOptions{
		Branch: branch_ref,
		Create: create,
	})

	if err != nil && err == plumbing.ErrReferenceNotFound {
		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: branch_ref,
			Create: true,
		})
	}

	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().StringVarP(&branch_name, "branch", "b", "", "The branch you want to checkout")
	checkoutCmd.Flags().BoolP("toggle", "t", false, "works exactly like git add, but does it for every repo")
}
