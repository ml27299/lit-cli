package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/ml27299/lit-cli/helpers"
	"github.com/ml27299/lit-cli/version"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates the lit cli",
	Long: `ex. lit update`,
	Run: updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
		
	var SilenceUsage = false
	for _, val := range args {
		if val == "silent" {
			SilenceUsage = true
		}
	}
	
	err := version.CheckForUpdate(SilenceUsage)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().String("foo", "", "A help for foo")
}
