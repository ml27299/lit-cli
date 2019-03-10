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

var (
	update_silent = false
	update_prompt = false	
)

func updateRun(cmd *cobra.Command, args []string) {
		
	for _, val := range args {
		if val == "silent" {
			update_silent = true
		}

		if val == "prompt" {
			update_prompt = true
		}
	}

	err := version.CheckForUpdate(update_silent, update_prompt)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolVarP(&update_prompt, "prompt", "p", false, "Will prompt if you want to update, not automatic")
	updateCmd.Flags().BoolVarP(&update_silent, "silent", "s", false, "Will not show any logs unless there is an update availiable")
}
