package bash
import (
	//"fmt"
	"os"
	"os/exec"
)

func Status(path string, args []string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    args = append([]string{"status"}, args...)
    cmd := exec.Command("git", args...)

    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	err = os.Chdir(current_path)
	if err != nil {
		return err
	}

	return nil
}

func StatusViaBash(path string, args string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    cmd := exec.Command("sh", "-c", "git status "+args)

    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	err = os.Chdir(current_path)
	if err != nil {
		return err
	}

	return nil
}