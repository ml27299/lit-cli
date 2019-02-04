package bash
import (
	"fmt"
	"os"
	"os/exec"
)

func Pull(path string, args []string) error {	
	current_path, err := os.Getwd()
	err = os.Chdir(path)

	if err != nil {
        return err
    }

    args = append([]string{"pull"}, args...)
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", out)
	err = os.Chdir(current_path)
	if err != nil {
		return err
	}

	return nil
}