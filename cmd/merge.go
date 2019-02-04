package cmd

import (
	//"fmt"
	//"os/exec"
	//"strings"
	"github.com/spf13/cobra"
	. "../helpers"
	"../helpers/paths"
	"../helpers/bash"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "merge files from a local or remote branch, analogous to \"git merge\"",
	Long: `ex. lit merge master`,
	Run: mergeRun,
	PostRun: linkRun,
}

func mergeRun(cmd *cobra.Command, args []string) {

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)

		Info("Entering "+*&status.Path+"...")

		err = bash.Merge(dir+"/"+*&status.Path, args)
    	CheckIfError(err)
	}

	Info("Entering /...")

	err = bash.Merge(dir+"/", args)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().BoolP("toggle", "t", false, "works exactly like git merge, but does it for every repo")
}
