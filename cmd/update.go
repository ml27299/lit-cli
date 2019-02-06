package cmd

import (
	"os"
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
	
	install_script_path := "https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh"
	args = append([]string{"checkout"}, args...)
   
	out, err := exec.Command("/bin/bash", "-c", "curl "+install_script_path).Output()
	CheckIfError(err)

	out_str := string(out[:])
	updatecmd := exec.Command("/bin/bash", "-c", out_str)

    updatecmd.Stdout = os.Stdout
	updatecmd.Stderr = os.Stderr

	err = updatecmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		CheckIfError(err)
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().String("foo", "", "A help for foo")
}
