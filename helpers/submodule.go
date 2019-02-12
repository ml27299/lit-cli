package helpers

import (
	"gopkg.in/src-d/go-git.v4"
	"errors"
	"os"
	"./paths"
	"./parser"
)

func GetSubmodules(path string) (git.Submodules, error){
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	submodules, err := worktree.Submodules()
	if err != nil {
		return nil, err
	}

	return submodules, nil
}

func FindSubmodule(submodules git.Submodules, value string) (*git.Submodule, error) {
	var response *git.Submodule

	dir, err := paths.FindRootDir()
	if err != nil {
		return response, err
	}

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		if err != nil{
			return response, err
 		}

		if status.Path == value {
			response = submodules[i]
			break
		}
	}


	if _, err := os.Stat(dir+"/.litconfig"); err == nil && response == nil {
		modules, err := parser.ConfigModules(dir+"/.litconfig")
		for i := 0; i < len(modules); i++ {
			if modules[i].Name == value {
				response, err = FindSubmodule(submodules, modules[i].Dest)
				if err != nil{
					return response, err
		 		}
				break
			}
		}
	}

	if response == nil {
		return response, errors.New("Couldn't find matching submodule for "+value)
	}

	return response, nil
}

func GetSubmoduleWorkTree(submodule *git.Submodule) (*git.Worktree, error){
	repo, err := submodule.Repository()
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	return worktree, nil
}