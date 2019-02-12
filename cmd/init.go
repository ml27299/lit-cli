package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"../helpers/parser"
	. "../helpers"
	"gopkg.in/src-d/go-git.v4"
	//"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"../helpers/paths"
	"../helpers/bash"
	"../helpers/resources"
	"github.com/mitchellh/go-homedir"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initializes a project, requires a lit.config.json",
	Long: `ex. lit init`,
	Run: initRun,
}

func initRun(cmd *cobra.Command, args []string) {

	var (
		home_dir string
		has_gitmodule = false
	)
	
	home_dir, err := homedir.Dir()
	home_dir, err = homedir.Expand(home_dir)
	CheckIfError(err)

	current_path, err := os.Getwd()
	CheckIfError(err)

    _, err = git.PlainInit(current_path, false)
    if err != nil && err != git.ErrRepositoryAlreadyExists {
    	CheckIfError(err)
    }

    root_dir, err := paths.FindRootDir()
    CheckIfError(err)

    submodules, err := GetSubmodules(root_dir)
	CheckIfError(err)

    _, err = os.Stat(root_dir+"/.gitmodules")
    if _, ok := err.(*os.PathError); err != nil && !ok {
    	CheckIfError(err)
    } else if err == nil {
    	has_gitmodule = true
    }

 //    if _, err := os.Stat(root_dir+"/lit.link.json"); os.IsNotExist(err) {

	// 	data, err := resources.Asset("lit.link.json")
	// 	file, err := os.Create(root_dir+"/lit.link.json")

	// 	defer file.Close()
		
	// 	CheckIfError(err)

	// 	_, err = file.Write(data)
	// 	CheckIfError(err)
	// }

	// if _, err := os.Stat(root_dir+"/lit.module.json"); os.IsNotExist(err) {

	// 	data, err := resources.Asset("lit.module.json")
	// 	file, err := os.Create(root_dir+"/lit.module.json")

	// 	defer file.Close()

	// 	CheckIfError(err)

	// 	_, err = file.Write(data)
	// 	CheckIfError(err)
	// }

	if _, err := os.Stat(root_dir+"/.gitignore"); os.IsNotExist(err) {

		data, err := resources.Asset(".gitignore")
		file, err := os.Create(root_dir+"/.gitignore")

		defer file.Close()

		CheckIfError(err)

		_, err = file.Write(data)
		CheckIfError(err)
	}

	if _, err := os.Stat(root_dir+"/.litconfig"); os.IsNotExist(err) {

		data, err := resources.Asset(".litconfig")
		file, err := os.Create(root_dir+"/.litconfig")

		defer file.Close()

		CheckIfError(err)

		_, err = file.Write(data)
		CheckIfError(err)
	}

    if has_gitmodule {

    	bash.SubmoduleUpdate()

  //   	auth, err := ssh.NewPublicKeysFromFile("git", home_dir+"/.ssh/id_rsa", "")
		// CheckIfError(err)

		// err = submodules.Update(&git.SubmoduleUpdateOptions{
		// 	Init: true,
		// 	RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		// 	Auth: auth,
		// })

		// CheckIfError(err)
    }

	info, err  := parser.Config()
	CheckIfError(err)

	var missing_submodules []parser.GitModule
	for _, gitmodule := range info.GitModules {
		found := false
		for _, submodule := range submodules {
			if submodule.Config().URL == gitmodule.Repo {
				found = true
				break
			}
		}
		if !found {
			missing_submodules = append(missing_submodules, gitmodule)
		}
	}

	for _, gitmodule := range missing_submodules {

		Info("Adding "+gitmodule.Repo)
		bash.SubmoduleAdd(gitmodule.Repo, gitmodule.Dest, gitmodule.Name)
    }
}

func init() {
	rootCmd.AddCommand(initCmd)
}
