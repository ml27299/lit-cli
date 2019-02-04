package cmd

import (
	//"fmt"
	"os"
	//"errors"
	"github.com/spf13/cobra"
	"../helpers/parser"
	. "../helpers"
	"../helpers/paths"
	"../helpers/prompt"
	"path/filepath"
	"strings"
	"gopkg.in/src-d/go-git.v4"
)

var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "creates a file and hard links it to the correspoding git module",
	Long: `ex. lit touch ./path/to/somefile.txt`,
	Run: touchRun,
}

func CreateAndLink(submodule *git.Submodule, info parser.ParseInfo, args []string) error {

	_, newfilename := filepath.Split(args[0])
	newfilepath, err := paths.Normalize(filepath.Dir(args[0]))
	submodule_path := submodule.Config().Path
	submodule_path, err = paths.Normalize(submodule_path)

	if err != nil {
		return err
	}
	
	matchedDestPaths, err := info.FindMatchingLinkItemsBySubmodule(submodule_path, newfilepath)
	if err != nil {
		return err
	}
	
	var index int
	if len(matchedDestPaths) == 1 {
		index = 0
	} else if len(matchedDestPaths) > 1 {
		index, err = prompt.PromtForDest(matchedDestPaths, newfilename, newfilepath, submodule_path)
		if err != nil {
			return err
		}
	}else if len(matchedDestPaths) == 0 {
		_continue, err := prompt.PromptForNoDest(newfilename, newfilepath)
		if !_continue {
			return nil
		}

		_, err = os.Create(args[0])
		if err != nil {
			return err
		}

		return nil
	}
	
	dest, err := paths.Normalize(matchedDestPaths[index])
	if err != nil {
		return err
	}

	newfilepath = strings.Replace(newfilepath, dest, "", -1)
	newfilepath = submodule_path+newfilepath

	_, err = os.Create(newfilepath+"/"+newfilename)
	og_newfilepath, err := paths.Normalize(filepath.Dir(args[0]))
	if err != nil {
		return err
	}

	err = Link(parser.Link{
		Dest: og_newfilepath+"/"+newfilename,
		Source: newfilepath+"/"+newfilename,
	})
	if err != nil {
		return err
	}

	return nil
}

func touchRun(cmd *cobra.Command, args []string) {
	
	_, newfilename := filepath.Split(args[0])
	newfilepath, err := paths.Normalize(filepath.Dir(args[0]))
	
	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	info, err := parser.Config()
	CheckIfError(err)
	
	items, err := info.FindMatchingLinkItems(newfilepath)
	CheckIfError(err)
	
	var submodule_paths []string
	for _, item := range items {
		paths := FindAssociateGitModulePaths(submodules, item.Sources...)
		submodule_paths = parser.AppendUnique(submodule_paths, paths...)
	}
	
	var matched_submodules []*git.Submodule
	for _, submodule_path := range submodule_paths {

		var _submodule *git.Submodule
		for _, submodule := range submodules {
			if submodule_path == submodule.Config().Path {
				_submodule = submodule
				break
			}
		}

		if _submodule != nil {
			matched_submodules = append(matched_submodules, _submodule)
		}
	}
	
	if len(matched_submodules) == 1 {

		err = CreateAndLink(matched_submodules[0], info, args)
		CheckIfError(err)

	} else if len(matched_submodules) > 1 {

		index, err := prompt.PromptForSubmodule(matched_submodules, newfilename, newfilepath)
		CheckIfError(err)

		err = CreateAndLink(matched_submodules[index], info, args)
		CheckIfError(err)

	}else if len(matched_submodules) == 0 {
		_, err = os.Create(args[0])
		CheckIfError(err)
	}
}

func init() {
	rootCmd.AddCommand(touchCmd)
}
