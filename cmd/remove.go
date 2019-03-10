package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/ml27299/helpers"
	"github.com/ml27299/helpers/paths"
	"github.com/ml27299/helpers/bash"
	"os/user"
	"os"
	"errors"
	"path/filepath"
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

	config_files, err := paths.FindConfig(dir)
	CheckIfError(err)

	current_path, err := os.Getwd()
	CheckIfError(err)

	for _, config_file := range config_files {
		config_file = filepath.Dir(config_file)
		if current_path == config_file {
			dir = config_file
			break
		}
	}

	err = os.Chdir(dir)
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	submodule, err := FindSubmodule(submodules, args[0])
	CheckIfError(err)

	// status, err := submodule.Status()
	// CheckIfError(err)

	Info("Removing "+*&submodule.Submodule.Config().Path+"...")
	err = bash.SubmoduleRemove(args[0], *&submodule.Submodule.Config().Path, submodule.Ext)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(removeCmd)
	//removeCmd.Flags().StringVarP(&removeName, "name", "n", "",  "The name of the submodule you want to remove")
	//removeCmd.MarkFlagRequired("name")
}
