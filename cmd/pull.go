package cmd

import (
	"fmt"
	//"os/exec"
	"os"
	//"strings"
	"errors"
	"gopkg.in/src-d/go-git.v4"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	. "../helpers"
	"../helpers/paths"
	"../helpers/bash"
	"github.com/mitchellh/go-homedir"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull files from a remote git repo, analogous to \"git pull\"",
	Long: `ex. lit pull origin master`,
	Run: pullRun,
	PostRun: linkRun,
}

func pullRun(cmd *cobra.Command, args []string) {

	home_dir, err := homedir.Dir()
	home_dir, err = homedir.Expand(home_dir)
	CheckIfError(err)

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	auth, err := ssh.NewPublicKeysFromFile("git", home_dir+"/.ssh/id_rsa", "")
	CheckIfError(err)
	
	var (
		ErrReferenceHasChanged = errors.New("reference has changed concurrently")
		remote_name = args[0]
		branch_ref = plumbing.NewBranchReferenceName(args[1])
	)

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)

		Info("Entering "+*&status.Path+"...")

		worktree, err := GetSubmoduleWorkTree(submodules[i])
		CheckIfError(err)
		
		pullerr := worktree.Pull(&git.PullOptions{
			RemoteName: remote_name,
			SingleBranch: true,
			ReferenceName: branch_ref,
			Auth: auth,
			Progress: os.Stdout,
		})

		if pullerr != nil && pullerr.Error() == "non-fast-forward update" {

			fmt.Println("non-fast-forward update")
			err := bash.Pull(dir+"/"+*&status.Path, args)
			CheckIfError(err)

		} else if pullerr != git.NoErrAlreadyUpToDate && pullerr != ErrReferenceHasChanged {
			CheckIfError(err)
		}
	}

	Info("Entering /...")
	repo, err := git.PlainOpen(dir)
	CheckIfError(err)

	worktree, err := repo.Worktree()
	CheckIfError(err)

	pullerr := worktree.Pull(&git.PullOptions{
		RemoteName: remote_name,
		SingleBranch: true,
		ReferenceName: branch_ref,
		Auth: auth,
		Progress: os.Stdout,
	})
	if pullerr != nil && pullerr.Error() == "non-fast-forward update" {
		
		err = bash.Pull(dir+"/", args)
		CheckIfError(err)

	} else if pullerr != nil {
		fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", pullerr))
	}
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().BoolP("toggle", "t", false, "works exactly like git pull, but does it for every repo")
}
