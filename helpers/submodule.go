package helpers

import (
	"gopkg.in/src-d/go-git.v4"
	"errors"
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