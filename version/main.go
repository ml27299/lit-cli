package version

import (
	"errors"
	"fmt"

	"../helpers/github"
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

func CheckForUpdate() error {
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

	fmt.Printf("Current Version: %s %s"+"\n", currentTag, currentPub)
	fmt.Printf("Latest Version: %s %s"+"\n", latestTag, latestPub)

	if latestTag > currentTag {
		fmt.Println("There is a more recent version of the Lit CLI available.")
		fmt.Println("curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh | sudo bash")
	} else {
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