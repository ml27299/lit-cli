package bash
import (
	"fmt"
	//"os"
	"os/exec"
)

func SubmoduleAdd(repo, dest string) error {
	out, err := exec.Command("git", "submodule", "add", repo, dest).Output()
	if err != nil {
        return err
    }

	if string(out[:]) != "" {
		fmt.Printf("%s\n", out)
	}

	return nil
}