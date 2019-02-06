package helpers

import (
	"gopkg.in/src-d/go-git.v4"
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