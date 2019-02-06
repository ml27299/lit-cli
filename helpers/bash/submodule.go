package bash
import (
	"os"
	"os/exec"
)

func SubmoduleAdd(repo, dest string) error {

	cmd := exec.Command("git", "submodule", "add", repo, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	return nil
}