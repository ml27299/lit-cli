package cmd

import (
	"path/filepath"
	"github.com/spf13/cobra"
	. "github.com/ml27299/lit-cli/helpers"
	"github.com/ml27299/lit-cli/helpers/paths"
	"github.com/ml27299/lit-cli/helpers/bash"
	"github.com/ml27299/lit-cli/helpers/parser"
	"os"
)

var branch string
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "initializes/adds submodules and hard links file(s)",
	Long: `ex. lit build`,
	Run: buildRun,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	updateRun(cmd, append(args, []string{"silent", "prompt"}...))
	// },
	PostRun: func(cmd *cobra.Command, args []string) {
		linkRun(cmd, args)
	},
}

func buildRun(cmd *cobra.Command, args []string) {

	checkout := func(dir string, submodules Modules) {
    	for i := 0; i < len(submodules); i++ {

			status, err := submodules[i].Status()
			CheckIfError(err)

			Info("Entering "+*&status.Path+"...")
			err = bash.Checkout(dir+"/"+*&status.Path, []string{branch})
			CheckIfError(err)
		}
    }

	dir, err := paths.FindRootDir()
	CheckIfError(err)
	
	err = os.Chdir(dir)
	CheckIfError(err)

	_, err = os.Stat(dir+"/.gitmodules")
    if _, ok := err.(*os.PathError); err != nil && !ok {
    	CheckIfError(err)
    } else if err == nil {
    	bash.SubmoduleUpdate()
    }

    submodules, err := GetSubmodules(dir)
	CheckIfError(err)

    config_files, err := paths.FindConfig(dir)
	CheckIfError(err)

	for _, config_file := range config_files {

		config_file_dir := filepath.Dir(config_file)
		err = os.Chdir(config_file_dir)
		CheckIfError(err)

		info, err := parser.ConfigViaPath(config_file_dir)
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

	    if len(missing_submodules) != 0 {
	    	bash.SubmoduleUpdate()
	    }
	}

	submodules, err = GetSubmodules(dir)
	CheckIfError(err)

	if branch != ""{
		checkout(dir, submodules)
	}else {
		Warning("Didnt supply a branch name (-b), be sure to use a concrete branch name or else the submodules within an application can get out of sync")
	}
}

func init() {
	buildCmd.Flags().StringVarP(&branch, "branch", "b", "", "")
	rootCmd.AddCommand(buildCmd)
}
