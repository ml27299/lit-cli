package helpers

import (
	//"fmt"
	"strings"
	"gopkg.in/src-d/go-git.v4"
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

func FindAssociateGitModulePaths(submodules []*git.Submodule, filepaths ...string) ([]string) {
	var response []string

	for _, filepath := range filepaths {
		for _, submodule := range submodules {
			config := submodule.Config()
			if strings.Contains(filepath, config.Path) {
				response = parser.AppendUnique(response, config.Path)
			}
		}
	}

	return response
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