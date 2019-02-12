package bash
import (
	"os"
	"os/exec"
)

func SubmoduleAdd(repo, dest string, name string) error {

	cmd := exec.Command("git", "submodule", "add", "--name "+name, repo, dest)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	return nil
}


func SubmoduleUpdate() error {

	cmd := exec.Command("git", "submodule", "update", "--init", "--recursive")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	return nil
}