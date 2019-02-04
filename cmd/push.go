package cmd

import (
	//"fmt"
	//"os/exec"
	"os"
	//"gopkg.in/src-d/go-git.v4/config"
	//"strings"
	"gopkg.in/src-d/go-git.v4"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	. "../helpers"
	"../helpers/paths"
	"github.com/mitchellh/go-homedir"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push files from a remote git repo, analogous to \"git push\"",
	Long: `ex. lit push origin master`,
	Run: pushRun,
}

func pushRun(cmd *cobra.Command, args []string) {

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
		remote_name = args[0]
		branch_ref = plumbing.NewBranchReferenceName(args[1])
	)

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)

		Info("Entering "+*&status.Path+"...")

		repo, err := submodules[i].Repository()
		CheckIfError(err)

		head, err := repo.Head()
		CheckIfError(err)

		ref := plumbing.NewHashReference(branch_ref, head.Hash())
		err = repo.Storer.SetReference(ref)
		CheckIfError(err)

		pusherr := repo.Push(&git.PushOptions{
			RemoteName: remote_name,
			Auth: auth,
			Progress: os.Stdout,
		})
		CheckIfError(pusherr)
	}

	Info("Entering /...")
	repo, err := git.PlainOpen(dir)
	CheckIfError(err)

	head, err := repo.Head()
	CheckIfError(err)

	ref := plumbing.NewHashReference(branch_ref, head.Hash())
	err = repo.Storer.SetReference(ref)
	CheckIfError(err)

	pusherr := repo.Push(&git.PushOptions{
		RemoteName: remote_name,
		Auth: auth,
		Progress: os.Stdout,
	})
	CheckIfError(pusherr)
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.Flags().BoolP("toggle", "t", false, "works exactly like git push, but does it for every repo")
}
