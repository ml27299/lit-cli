package bash
import (
	"fmt"
	"os"
	"os/exec"
)

func Merge(path string, args []string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    args = append([]string{"merge"}, args...)
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return err
	}

	if string(out[:]) == "Already up to date" {
		fmt.Printf("%s\n", out)
	}

	err = os.Chdir(current_path)
	if err != nil {
		return err
	}

	return nil
}