package bash
import (
	"os"
	"os/exec"
	"github.com/ml27299/helpers/paths"
	//"fmt"
)

func SubmoduleAdd(repo, dest string, name string) error {

	cmd := exec.Command("git", "submodule", "add", "-f", "--name", name, repo, dest)
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

func SubmoduleRemove(name, path string, ext string) error {

	dir, err := paths.FindRootDir()
	if err != nil {
		return err
	}

	cmd_str := "git config -f .gitmodules --remove-section submodule."+name
	if ext != "" {
		cmd_str = "git config -f "+dir+"/"+ext+"/.gitmodules --remove-section submodule."+name
	}
	cmd := exec.Command("sh", "-c", cmd_str)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
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

	cmd_str = "git config -f .git/config --remove-section submodule."+name
	if ext != "" {
		cmd_str = "git config -f "+dir+"/.git/modules/"+ext+"/config --remove-section submodule."+name
	}
	cmd = exec.Command("sh", "-c", cmd_str)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "rm", "-rf", "--cached", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	cmd_str = ".git/modules/"+name
	if ext != "" {
		cmd_str = dir+"/.git/modules/"+ext+"/modules/"+name
	}
	cmd = exec.Command("rm", "-rf", cmd_str)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	if "./"+path == "./" {
		return nil
	}

	cmd = exec.Command("rm", "-rf", "./"+path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	return nil
}