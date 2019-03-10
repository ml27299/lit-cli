package version

import (
	"os"
	"os/exec"
	"errors"
	"fmt"

	"github.com/ml27299/lit-cli/helpers/github"
	"github.com/ml27299/lit-cli/helpers/prompt"
)

var (
	CurrVersion string
	CurrCommit  string
	api         = github.NewGithubClient()
)

func PrintVersion() error {
	version := CurrVersion
	gitCommit := CurrCommit
	
	if !isValidVersion(version) {
		return errors.New("Not a correct cli version")
	}

	fmt.Printf("Current Version: %s"+"\n", version)
	fmt.Printf("Current Commit: %s"+"\n", gitCommit)
	return nil
}

func RunLatestUpdate() error {

	install_script_path := "https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh"
	
	cmd := exec.Command("sh", "-c", "curl -s "+install_script_path+" | sudo bash")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func CheckForUpdate(silent bool) error {
	version := CurrVersion

	if !isValidVersion(version) {
		fmt.Println("Current version not tagged, please reinstall...")
		fmt.Println("curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh | sudo bash")
		return nil
	}

	latestTagResp, err := api.RepoLatestRequest("ml27299", "lit-cli")
	if err != nil {
		fmt.Println(err)
		latestTagResp.TagName = "N/A"
	}

	currentTagResp, err := api.RepoTagRequest("ml27299", "lit-cli", string("v")+version)
	if err != nil {
		fmt.Println("Release info not found, please upgrade.")
		fmt.Println("curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh | sudo bash")
		return nil
	}

	currentPub := currentTagResp.PublishedAt.Format("2006.01.02")
	currentTag := currentTagResp.TagName
	latestPub := latestTagResp.PublishedAt.Format("2006.01.02")
	latestTag := latestTagResp.TagName

	if !silent {
		fmt.Printf("Current Version: %s %s"+"\n", currentTag, currentPub)
		fmt.Printf("Latest Version: %s %s"+"\n", latestTag, latestPub)
	}

	if latestTag > currentTag {

		if silent {
			fmt.Printf("Current Version: %s %s"+"\n", currentTag, currentPub)
			fmt.Printf("Latest Version: %s %s"+"\n", latestTag, latestPub)
		}

		fmt.Println("There is a more recent version of the Lit CLI available.")
		str, err := prompt.PromptForUpdate()
		if err != nil {
			return err
		}

		if str == "yes" || str == "y" {
			err = RunLatestUpdate()
			if err != nil {
				return err
			}
		}else {
			fmt.Println("Its recomended to be up to date, you can always manually update with the command below")
			fmt.Println("curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh | sudo bash")
		}

	} else if !silent {
		fmt.Println("Running latest version")
	}

	return nil
}

func isValidVersion(version string) bool {
	if len(version) == 0 {
		return false
	}
	return true
}