package bash
import (
	"fmt"
	"os"
	"os/exec"
)

func hasChanges() (bool, error) {
	has_changes, err := exec.Command("git", "diff", "--exit-code").Output()
	if err != nil && err.Error() != "exit status 1" {
		return false, err
	}

	if string(has_changes[:]) == "" {
		has_changes, err = exec.Command("git", "diff", "--cache", "--exit-code").Output()
		if err != nil && err.Error() != "exit status 1" {
			return false, err
		}
	}

	if string(has_changes[:]) == "" {

		untracked_files, err := exec.Command("git", "ls-files", "--other", "--exclude-standard", "--directory").Output()
		if err != nil && err.Error() != "exit status 1" {
			return false, err
		}

		if string(untracked_files[:]) == "" {
			return false, nil
		}

		fmt.Println("Untacked files...")
		fmt.Println(string(untracked_files))
		return false, nil
	}

	return true, nil
}

func Commit(path string, args []string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    changes, err := hasChanges()
    if err != nil {
		return err
	}

	if !changes {
		return nil
	}

    args = append([]string{"commit"}, args...)
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

func CommitViaBash(path string, args string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    changes, err := hasChanges()
    if err != nil {
		return err
	}

	if !changes {
		return nil
	}

    cmd := exec.Command("sh", "-c", "git commit "+args)

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