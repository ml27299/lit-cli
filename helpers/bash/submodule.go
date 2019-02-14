package bash
import (
	"os"
	"os/exec"
	//"fmt"
)

func SubmoduleAdd(repo, dest string, name string) error {

	cmd := exec.Command("git", "submodule", "add", "--name", name, repo, dest)
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

// eval "git config -f .gitmodules --remove-section submodule.$SUBMODULE > /dev/null"
// 	eval "git add .gitmodules"

// 	eval "git config -f .git/config --remove-section submodule.$SUBMODULE > /dev/null"
// 	eval "git rm -f --cached $SUBMODULE"

// 	eval "rm -rf .git/modules/$SUBMODULE"

func SubmoduleRemove(name, path string) error {

	cmd := exec.Command("sh", "-c", "git config -f .gitmodules --remove-section submodule."+name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	cmd = exec.Command("sh", "-c", "git add .gitmodules")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	cmd = exec.Command("sh", "-c", "git config -f .git/config --remove-section submodule."+name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "rm", "-f", "--cached", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	cmd = exec.Command("rm", "-rf", ".git/modules/"+name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	return nil
}