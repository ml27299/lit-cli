package cmd

import (
	"os"
	"github.com/spf13/cobra"
	. "github.com/ml27299/lit-cli/helpers"
	"gopkg.in/src-d/go-git.v4"
	"github.com/ml27299/lit-cli/helpers/paths"
	"github.com/ml27299/lit-cli/helpers/resources"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a project, requires a lit.config.json",
	Long: `ex. lit init`,
	Run: initRun,
	PostRun: func(cmd *cobra.Command, args []string) {
		updateRun(cmd, append(args, "silent"))
	},
}

func initRun(cmd *cobra.Command, args []string) {

	current_path, err := os.Getwd()
	CheckIfError(err)

    _, err = git.PlainInit(current_path, false)
    if err != nil && err != git.ErrRepositoryAlreadyExists {
    	CheckIfError(err)
    }

    root_dir, err := paths.FindRootDir()
    CheckIfError(err)

	if _, err := os.Stat(root_dir+"/.gitignore"); os.IsNotExist(err) {

		data, err := resources.Asset(".gitignore")
		file, err := os.Create(root_dir+"/.gitignore")

		defer file.Close()

		CheckIfError(err)

		_, err = file.Write(data)
		CheckIfError(err)
	}

	if _, err := os.Stat(root_dir+"/.litconfig"); os.IsNotExist(err) {

		data, err := resources.Asset(".litconfig")
		file, err := os.Create(root_dir+"/.litconfig")

		defer file.Close()

		CheckIfError(err)

		_, err = file.Write(data)
		CheckIfError(err)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
