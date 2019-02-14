package cmd

import (
	"github.com/spf13/cobra"
	. "../helpers"
	"../helpers/paths"
	"../helpers/bash"
	"../helpers/parser"
	"os"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "initializes/adds submodules and hard links file(s)",
	Long: `ex. lit build`,
	Run: buildRun,
	PostRun: linkRun,
}

func buildRun(cmd *cobra.Command, args []string) {

	dir, err := paths.FindRootDir()
	CheckIfError(err)
	
	err = os.Chdir(dir)
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	_, err = os.Stat(dir+"/.gitmodules")
    if _, ok := err.(*os.PathError); err != nil && !ok {
    	CheckIfError(err)
    } else if err == nil {
    	bash.SubmoduleUpdate()
    }

	info, err  := parser.Config()
	CheckIfError(err)

	var missing_submodules []parser.GitModule
	for _, gitmodule := range info.GitModules {
		found := false
		for _, submodule := range submodules {
			if submodule.Config().URL == gitmodule.Repo {
				found = true
				break
			}
		}
		if !found {
			missing_submodules = append(missing_submodules, gitmodule)
		}
	}

	for _, gitmodule := range missing_submodules {

		Info("Adding "+gitmodule.Repo)
		bash.SubmoduleAdd(gitmodule.Repo, gitmodule.Dest, gitmodule.Name)
    }
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
