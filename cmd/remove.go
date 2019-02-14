package cmd

import (
	"github.com/spf13/cobra"
	. "../helpers"
	"../helpers/paths"
	"../helpers/bash"
	"os/user"
	"errors"
)

// var (
// 	removeName string
// )

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes a submodule from a project",
	Long: `ex. lit remove {{submodule.name || submodule.path}}`,
	Run: removeRun,
}

func removeRun(cmd *cobra.Command, args []string) {

	User, err := user.Current()
	CheckIfError(err)

	if User.Uid != "0" {
		CheckIfError(errors.New("Please run as root"))
	}

	dir, err := paths.FindRootDir()
	CheckIfError(err)
	
	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	submodule, err := FindSubmodule(submodules, args[0])
	CheckIfError(err)

	status, err := submodule.Status()
	CheckIfError(err)

	Info("Removing "+*&status.Path+"...")
	err = bash.SubmoduleRemove(args[0], *&status.Path)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(removeCmd)
	//removeCmd.Flags().StringVarP(&removeName, "name", "n", "",  "The name of the submodule you want to remove")
	//removeCmd.MarkFlagRequired("name")
}
