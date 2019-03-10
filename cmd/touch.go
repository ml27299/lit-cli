package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/ml27299/lit-cli/helpers/parser"
	. "github.com/ml27299/lit-cli/helpers"
	"github.com/ml27299/lit-cli/helpers/paths"
	"github.com/ml27299/lit-cli/helpers/prompt"
	"path/filepath"
	"strings"
)

var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "creates a file and hard links it to the correspoding git module",
	Long: `ex. lit touch ./path/to/somefile.txt`,
	Run: touchRun,
	PostRun: func(cmd *cobra.Command, args []string) {
		updateRun(cmd, append(args, "silent"))
	},
}

func CreateAndLink(source string, items []parser.LinkItem, args []string) error {

	_, newfilename := filepath.Split(args[0])
	newfilepath, err := paths.Normalize(filepath.Dir(args[0]))
	source_path, err := paths.Normalize(filepath.Dir(source))

	if err != nil {
		return err
	}	
	
	var index int
	if len(items) == 1 {
		index = 0
	} else if len(items) > 1 {

		var destPaths []string
		for _, item := range items { 
			destPaths = append(destPaths, item.Dest)
		}

		index, err = prompt.PromtForDest(destPaths, newfilename, newfilepath, source_path)
		if err != nil {
			return err
		}
	}else if len(items) == 0 {
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
	
	dest, err := paths.Normalize(items[index].Dest)
	if err != nil {
		return err
	}

	newfilepath = strings.Replace(newfilepath, dest, "", -1)
	newfilepath = source_path+newfilepath

	_, err = os.Create(newfilepath+"/"+newfilename)
	og_newfilepath, err := paths.Normalize(filepath.Dir(args[0]))
	if err != nil {
		return err
	}

	Link(parser.Link{
		Dest: og_newfilepath+"/"+newfilename,
		Source: newfilepath+"/"+newfilename,
	})

	return nil
}

func touchRun(cmd *cobra.Command, args []string) {
	dir, err := paths.FindRootDir()
	CheckIfError(err)

	_, newfilename := filepath.Split(args[0])
	newfilepath, err := paths.Normalize(filepath.Dir(args[0]))

	config_files, err := paths.FindConfig(dir)
	CheckIfError(err)

	for _, config_file := range config_files {

		config_file_dir := filepath.Dir(config_file)
		err = os.Chdir(config_file_dir)
		CheckIfError(err)

		info, err := parser.ConfigViaPath(config_file_dir)
		CheckIfError(err)
	
		items, err := info.FindMatchingLinkItems(newfilepath)
		CheckIfError(err)
		
		var (
			matched_sources []string
			matched_items []parser.LinkItem
		)

		for _, item := range items {

			sources, err := item.FindMatchingSources(newfilepath)
			CheckIfError(err)
			
			matched_sources = parser.AppendUnique(matched_sources, sources...)
			if len(sources) > 0 {
				matched_items = append(matched_items, item)
			}
		}
		
		if len(matched_sources) == 0 && len(matched_items) == 0 {

			_, err = os.Create(args[0])
			CheckIfError(err)

		} else if len(matched_sources) == 1 {

			err = CreateAndLink(matched_sources[0], matched_items, args)
			CheckIfError(err)

		} else if len(matched_sources) > 1 {

			index, err := prompt.PromptForMultiSource(matched_sources, newfilename, newfilepath)
			CheckIfError(err)

			err = CreateAndLink(matched_sources[index], matched_items, args)
			CheckIfError(err)
		}
	}
}

func init() {
	rootCmd.AddCommand(touchCmd)
}
