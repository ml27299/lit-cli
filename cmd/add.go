package cmd

import (
	"path/filepath"
	"os"
	"github.com/spf13/cobra"
	. "github.com/ml27299/lit-cli/helpers"
	"github.com/ml27299/lit-cli/helpers/paths"
	"github.com/ml27299/lit-cli/helpers/bash"
	"github.com/ml27299/lit-cli/helpers/parser"
	Args "github.com/ml27299/lit-cli/helpers/args"
	"github.com/ml27299/lit-cli/helpers/prompt"
)

var (
	addSlug = "git-add"
	addStringArgs [1]string
	addBoolArgs [17]bool
	addStringArgIndexMap = make(map[int]Args.StringArg)
	addBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: DocRoot+"/"+addSlug,
	Long: `ex. lit add .`,
	Run: addRun,
	PostRun: func(cmd *cobra.Command, args []string) {
		updateRun(cmd, args)
	},
}

func addRun(cmd *cobra.Command, args []string) {
	
	add := func(dir string, submodules Modules) {
    	for i := 0; i < len(submodules); i++ {

			status, err := submodules[i].Status()
			CheckIfError(err)

			if interactive {
				command, err := prompt.PromptForInteractive(args, submodules[i])
				CheckIfError(err)

				err = bash.AddViaBash(dir+"/"+*&status.Path, command)
				CheckIfError(err)

				continue
			}

			Info("Entering "+*&status.Path+"...")
			err = bash.Add(dir+"/"+*&status.Path, args)
			CheckIfError(err)
		}
    }

	for index, arg := range addBoolArgIndexMap {
		addBoolArgIndexMap[index] = arg.SetValue(addBoolArgs[index])
	}
	for index, arg := range addStringArgIndexMap {
		addStringArgIndexMap[index] = arg.SetValue(addStringArgs[index])
	}

	_args := Args.GenerateCommand(addStringArgIndexMap, addBoolArgIndexMap)
	args = append(_args, args...)

	dir, err := paths.FindRootDir()
	CheckIfError(err)
	
	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	config_files, err := paths.FindConfig(dir)
	CheckIfError(err)

	for _, config_file := range config_files {

		config_file_dir := filepath.Dir(config_file)
		err = os.Chdir(config_file_dir)
		CheckIfError(err)

		info, err := parser.ConfigViaPath(config_file_dir)
		CheckIfError(err)

		links, err := info.GetLinks()
		CheckIfError(err)

		err = UpdateGitignore(config_file_dir, links)
		CheckIfError(err)
	}

	if submodule != "" {
		
		_submodule, err := FindSubmodule(submodules, submodule)
		CheckIfError(err)

		status, err := _submodule.Status()
		CheckIfError(err)

		submodules, err = GetSubmodules(dir+"/"+*&status.Path)
		add(dir, submodules)

		Info("Entering "+*&status.Path+"...")
		err = bash.Add(dir+"/"+*&status.Path, args)
		CheckIfError(err)

		return 
	}

	add(dir, submodules)

	Info("Entering /...")
	err = bash.Add(dir, args)
	CheckIfError(err)
}




func init() {
	rootCmd.AddCommand(addCmd)
	
	addStringArgIndexMap[0] = Args.StringArg{ Long: "chmod", Short: "" } 
	for index, val := range addStringArgIndexMap {
		addCmd.Flags().StringVarP(&addStringArgs[index], val.Long, val.Short, "", DocRoot+"/"+addSlug+"#"+addSlug+"-"+val.Long)
	}

	addBoolArgIndexMap[0] = Args.BoolArg{ Long: "dry-run", Short: "n", } 
	addBoolArgIndexMap[1] = Args.BoolArg{ Long: "verbose", Short: "v" } 
	addBoolArgIndexMap[2] = Args.BoolArg{ Long: "force", Short: "f" } 
	addBoolArgIndexMap[3] = Args.BoolArg{ Long: "interactive", Short: "i" } 
	addBoolArgIndexMap[4] = Args.BoolArg{ Long: "patch", Short: "p" } 
	addBoolArgIndexMap[5] = Args.BoolArg{ Long: "edit", Short: "e" } 
	addBoolArgIndexMap[6] = Args.BoolArg{ Long: "update", Short: "u" } 
	addBoolArgIndexMap[7] = Args.BoolArg{ Long: "all", Short: "A" } 
	addBoolArgIndexMap[8] = Args.BoolArg{ Long: "no-ignore-removal", Short: "" } 
	addBoolArgIndexMap[9] = Args.BoolArg{ Long: "no-all", Short: "" } 
	addBoolArgIndexMap[10] = Args.BoolArg{ Long: "ignore-removal", Short: "" } 
	addBoolArgIndexMap[11] = Args.BoolArg{ Long: "intent-to-add", Short: "N" } 
	addBoolArgIndexMap[12] = Args.BoolArg{ Long: "refresh", Short: "" } 
	addBoolArgIndexMap[13] = Args.BoolArg{ Long: "ignore-errors", Short: "" } 
	addBoolArgIndexMap[14] = Args.BoolArg{ Long: "ignore-missing", Short: "" } 
	addBoolArgIndexMap[15] = Args.BoolArg{ Long: "no-warn-embedded-repo", Short: "" } 
	addBoolArgIndexMap[16] = Args.BoolArg{ Long: "renormalize", Short: "" } 

	for index, val := range addBoolArgIndexMap {
		addCmd.Flags().BoolVarP(&addBoolArgs[index], val.Long, val.Short, false, DocRoot+"/"+addSlug+"#"+addSlug+"-"+val.Long)
	}
}
