package cmd

import (
	"github.com/ml27299/lit-cli/version"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Lit CLI version",
		Long:  "The lit-cli semantic version and git commit tied to that release.",
		RunE:  printVersion,
	}

	upgradeCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "Check for newer version of Lit CLI",
		Long:  "Check for newer version of Lit CLI",
		RunE:  upgradeCheck,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(upgradeCmd)
}

func printVersion(cmd *cobra.Command, args []string) error {
	// Silence Usage as we have now validated command input
	cmd.SilenceUsage = true
	err := version.PrintVersion()
	if err != nil {
		return err
	}
	return nil
}

func upgradeCheck(cmd *cobra.Command, args []string) error {
	// Silence Usage as we have now validated command input
	cmd.SilenceUsage = true

	err := version.CheckForUpdate()
	if err != nil {
		return err
	}
	return nil
}