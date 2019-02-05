package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
	. "../helpers"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates the lit cli",
	Long: `ex. lit update`,
	Run: updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
	out, err := exec.Command("/bin/bash", "-c", "curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh").Output()
	CheckIfError(err)

	//fmt.Printf("%s\n", out)

	out_str := string(out[:])
	out, err = exec.Command("/bin/bash", "-c", out_str).Output()
	CheckIfError(err)

	fmt.Printf("%s\n", out)
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().String("foo", "", "A help for foo")
}
