package bash
import (
	"os"
	"os/exec"
)

func Add(path string, args []string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    args = append([]string{"add"}, args...)
    cmd := exec.Command("git", args...)

    cmd.Stdin = os.Stdin
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

func AddViaBash(path string, args string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    cmd := exec.Command("sh", "-c", "git add "+args)

    cmd.Stdin = os.Stdin
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