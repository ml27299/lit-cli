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
	commitSlug = "git-commit"
	commitStringArgs [12]string
	commitBoolArgs [25]bool
	commitStringArgIndexMap = make(map[int]Args.StringArg)
	commitBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short:  DocRoot+"/"+commitSlug,
	Long: `ex. lit commit -am "my commit message"`,
	Run: commitRun,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	updateRun(cmd, append(args, []string{"silent", "prompt"}...))
	// },
}

func commitRun(cmd *cobra.Command, args []string) {
	//var commited_status_paths []string
	commit := func(dir string, submodules Modules) {
    	for i := 0; i < len(submodules); i++ {

			status, err := submodules[i].Status()
			CheckIfError(err)

			changes, err := bash.HasCommitChanges(dir+"/"+*&status.Path)
			CheckIfError(err)
			if !changes {
				Info("Entering "+*&status.Path+"...")
				continue
			}

			if interactive {
				command, err := prompt.PromptForInteractive(args, submodules[i])
				CheckIfError(err)

				err = bash.CommitViaBash(dir+"/"+*&status.Path, command)
				CheckIfError(err)

				continue
			}

			Info("Entering "+*&status.Path+"...")
			err = bash.Commit(dir+"/"+*&status.Path, args)
			CheckIfError(err)

			//commited_status_paths = append(commited_status_paths, status.Path)
		}
    }
   
	for index, arg := range commitBoolArgIndexMap {
		commitBoolArgIndexMap[index] = arg.SetValue(commitBoolArgs[index])
	}
	for index, arg := range commitStringArgIndexMap {
		commitStringArgIndexMap[index] = arg.SetValue(commitStringArgs[index])
	}
	
	_args := Args.GenerateCommand(commitStringArgIndexMap, commitBoolArgIndexMap)
	args = append(_args, args...)

	dir, err := paths.FindRootDir()
	CheckIfError(err)
	
	submodules, err := GetSubmodules(dir, dir)
	CheckIfError(err)
	
	config_files, err := paths.FindConfig(dir)
	CheckIfError(err)
	
	for _, config_file := range config_files {
		
		config_file_dir := filepath.Dir(config_file)
		err = os.Chdir(config_file_dir)
		CheckIfError(err)
		
		info, err := parser.ConfigViaPath(config_file_dir, dir)
		CheckIfError(err)
		
		links, err := info.GetLinks(dir)
		CheckIfError(err)
		
		err = UpdateGitignore(config_file_dir, links)
		CheckIfError(err)
	}

	if submodule != "" {
		
		_submodule, err := FindSubmodule(submodules, submodule)
		CheckIfError(err)

		status, err := _submodule.Status()
		CheckIfError(err)

		submodules, err = GetSubmodules(dir+"/"+*&status.Path, dir)
		commit(dir, submodules)

		err = bash.Commit(dir+"/"+*&status.Path, args)
		CheckIfError(err)

		return 
	}

	Info("Entering /...")
	err = bash.Commit(dir, args)
	CheckIfError(err)

	commit(dir, submodules)
	SyncCommitIds(submodules, dir)
	
	
	// if len(commited_status_paths) > 0 {

	// 	Info("Entering /...")
	// 	for _, commited_status_path := range commited_status_paths {
	// 		bash.AddViaBash(dir, commited_status_path)
	// 	}

	// 	err = bash.CommitViaBash(dir, "-m \"synced commit id\"")
	// 	CheckIfError(err)
	// }
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitStringArgIndexMap[0] = Args.StringArg{ Long: "message", Short: "m" } 
	commitStringArgIndexMap[1] = Args.StringArg{ Long: "reuse-message", Short: "C" } 
	commitStringArgIndexMap[2] = Args.StringArg{ Long: "reedit-message", Short: "c" } 
	commitStringArgIndexMap[3] = Args.StringArg{ Long: "fixup", Short: "" } 
	commitStringArgIndexMap[4] = Args.StringArg{ Long: "squash", Short: "" } 
	commitStringArgIndexMap[5] = Args.StringArg{ Long: "file", Short: "" } 
	commitStringArgIndexMap[6] = Args.StringArg{ Long: "author", Short: "" } 
	commitStringArgIndexMap[7] = Args.StringArg{ Long: "date", Short: "" } 
	commitStringArgIndexMap[8] = Args.StringArg{ Long: "template", Short: "t" } 
	commitStringArgIndexMap[9] = Args.StringArg{ Long: "cleanup", Short: "" } 
	commitStringArgIndexMap[10] = Args.StringArg{ Long: "untracked-files", Short: "u" } 
	commitStringArgIndexMap[11] = Args.StringArg{ Long: "gpg-sign", Short: "S" } 

	
	for index, val := range commitStringArgIndexMap {
		commitCmd.Flags().StringVarP(&commitStringArgs[index], val.Long, val.Short, "",  DocRoot+"/"+commitSlug+"#"+commitSlug+"-"+val.Long)
	}

	commitBoolArgIndexMap[0] = Args.BoolArg{ Long: "all", Short: "a" } 
	commitBoolArgIndexMap[1] = Args.BoolArg{ Long: "patch", Short: "p" } 
	commitBoolArgIndexMap[2] = Args.BoolArg{ Long: "reset-author", Short: "" } 
	commitBoolArgIndexMap[3] = Args.BoolArg{ Long: "short", Short: "" } 
	commitBoolArgIndexMap[4] = Args.BoolArg{ Long: "branch", Short: "" } 
	commitBoolArgIndexMap[5] = Args.BoolArg{ Long: "porcelain", Short: "" } 
	commitBoolArgIndexMap[6] = Args.BoolArg{ Long: "long", Short: "" } 
	commitBoolArgIndexMap[7] = Args.BoolArg{ Long: "null", Short: "z" } 
	commitBoolArgIndexMap[8] = Args.BoolArg{ Long: "signoff", Short: "s" } 
	commitBoolArgIndexMap[9] = Args.BoolArg{ Long: "no-verify", Short: "n" } 
	commitBoolArgIndexMap[10] = Args.BoolArg{ Long: "allow-empty", Short: "" } 
	commitBoolArgIndexMap[11] = Args.BoolArg{ Long: "allow-empty-message", Short: "" } 
	commitBoolArgIndexMap[12] = Args.BoolArg{ Long: "edit", Short: "e" } 
	commitBoolArgIndexMap[13] = Args.BoolArg{ Long: "no-edit", Short: "" } 
	commitBoolArgIndexMap[14] = Args.BoolArg{ Long: "amend", Short: "" } 
	commitBoolArgIndexMap[15] = Args.BoolArg{ Long: "no-post-rewrite", Short: "" } 
	commitBoolArgIndexMap[16] = Args.BoolArg{ Long: "include", Short: "i" } 
	commitBoolArgIndexMap[17] = Args.BoolArg{ Long: "only", Short: "o" } 
	commitBoolArgIndexMap[18] = Args.BoolArg{ Long: "verbose", Short: "v" } 
	commitBoolArgIndexMap[19] = Args.BoolArg{ Long: "quiet", Short: "q" } 
	commitBoolArgIndexMap[20] = Args.BoolArg{ Long: "dry-run", Short: "" } 
	commitBoolArgIndexMap[21] = Args.BoolArg{ Long: "status", Short: "" } 
	commitBoolArgIndexMap[22] = Args.BoolArg{ Long: "no-status", Short: "" }
	commitBoolArgIndexMap[23] = Args.BoolArg{ Long: "no-gpg-sign", Short: "" } 
	commitBoolArgIndexMap[24] = Args.BoolArg{ Long: "interactive", Short: "" } 
	
	for index, val := range commitBoolArgIndexMap {
		commitCmd.Flags().BoolVarP(&commitBoolArgs[index], val.Long, val.Short, false,  DocRoot+"/"+commitSlug+"#"+commitSlug+"-"+val.Long)
	}
}
