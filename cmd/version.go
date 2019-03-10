package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/ml27299/lit-cli/helpers"
	"github.com/ml27299/lit-cli/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Lit CLI version",
	Long:  "The lit-cli semantic version and git commit tied to that release.",
	Run: versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
		
	cmd.SilenceUsage = true
	err := version.PrintVersion()

	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}