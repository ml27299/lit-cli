package helpers

import (
	"gopkg.in/src-d/go-git.v4"
	"errors"
	"os"
	"github.com/ml27299/helpers/paths"
	"github.com/ml27299/helpers/parser"
	"strings"
	//"fmt"
	"os/exec"
	"path/filepath"
)

type ModuleConfig struct{
	Name string
    Path string
    URL string
    Branch string
}

type ModuleStatus struct{
	Path string
}

type Module struct{
	Ext string
	Submodule *git.Submodule
}

type Modules []*Module
func (m *Module) Status() (*ModuleStatus, error) {
	submodule_conf := m.Submodule.Config()
	if m.Ext != "" {
		return &ModuleStatus {
			Path: m.Ext+"/"+submodule_conf.Path,
		}, nil
	}else {
		return &ModuleStatus {
			Path: submodule_conf.Path,
		}, nil
	}
}
func (m *Module) Config() (*ModuleConfig) {
	submodule_conf := m.Submodule.Config()
	if m.Ext != "" {
		return &ModuleConfig {
			Path: m.Ext+"/"+submodule_conf.Path,
			Name: submodule_conf.Name,
			URL: submodule_conf.URL,
			Branch: submodule_conf.Branch,
		}
	}else {
		return &ModuleConfig {
			Path: submodule_conf.Path,
			Name: submodule_conf.Name,
			URL: submodule_conf.URL,
			Branch: submodule_conf.Branch,
		}
	}
}

func GetSubmodules(path string) (Modules, error){

	var response Modules
	repo, err := git.PlainOpen(path)
	if err != nil {
		return response, err
	}

	root_dir, err := paths.FindRootDir()
	if err != nil {
		return response, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return response, err
	}

	submodules, err := worktree.Submodules()
	if err != nil {
		return response, err
	}

	for _, submodule := range submodules {

		submodule_config := submodule.Config()
		ext := ""
		if path != root_dir {
			ext = strings.Replace(path, root_dir+"/", "", 1)
		}

		response = append(response, &Module{
			Ext: ext,
			Submodule: submodule,
		})

		modules, err := GetSubmodules(path+"/"+submodule_config.Path)
		if err != nil {
			return response, err
		}

		response = append(response, modules...)
	}

	return response, nil
}

// func FindSubmoduleAndName(submodules git.Submodules, value string) (*git.Submodule, string, error) {
// 	var (
// 		submodule *git.Submodule
// 		name string
// 	)

// 	if _, err := os.Stat(dir+"/.litconfig"); err == nil && response == nil {
// 		modules, err := parser.ConfigModules(dir+"/.litconfig")
// 		for i := 0; i < len(modules); i++ {
// 			if modules[i].Name == value {
				
// 				submodule, err = FindSubmodule(submodules, modules[i].Dest)
// 				if err != nil{
// 					return submodule, name, err
// 		 		}

// 				break
// 			}
// 		}
// 	}else {

// 	}

// 	return submodule, name, nil
// }

func FindSubmodule(submodules Modules, value string) (*Module, error) {
	var response *Module

	dir, err := paths.FindRootDir()
	if err != nil {
		return response, err
	}

	config_files, err := paths.FindConfig(dir)
	if err != nil {
		return response, err
	}

	current_path, err := os.Getwd()
	if err != nil {
		return response, err
	}
	
	var config_file_match string
	for _, config_file := range config_files {
		config_file = filepath.Dir(config_file)
		if current_path == config_file {
			config_file_match = config_file
			break
		}
	}

	ext := ""
	if config_file_match != "" && config_file_match != dir {
		ext = strings.Replace(config_file_match, dir+"/", "", 1)
	}
	
	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		if err != nil{
			return response, err
 		}

		if ext != "" && status.Path == ext+"/"+value {
			response = submodules[i]
			break
		}else if status.Path == value {
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

		cmd := "git config -f .gitmodules --list"
		if ext != "" {
			cmd = "git config -f "+dir+"/"+ext+"/.gitmodules --list"
		}

		out, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil && err.Error() != "exit status 1" {
			return response, err
		}

		split := strings.Split(string(out), "\n")
		for _, item := range split {
			if strings.Contains(item, "submodule."+value+".path") {
				path := strings.Split(item, "submodule."+value+".path=")
				response, err = FindSubmodule(submodules, path[len(path) - 1])
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