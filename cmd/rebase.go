package cmd

import (
	"github.com/spf13/cobra"
	. "github.com/ml27299/lit-cli/helpers"
	"github.com/ml27299/lit-cli/helpers/paths"
	"github.com/ml27299/lit-cli/helpers/bash"
	Args "github.com/ml27299/lit-cli/helpers/args"
	"github.com/ml27299/lit-cli/helpers/prompt"
)

var (
	rebaseSlug = "git-rebase"
	rebaseStringArgs [8]string
	rebaseBoolArgs [28]bool
	rebaseStringArgIndexMap = make(map[int]Args.StringArg)
	rebaseBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short:  DocRoot+"/"+rebaseSlug,
	Long: `ex. lit rebase`,
	Run: rebaseRun,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	updateRun(cmd, append(args, []string{"silent", "prompt"}...))
	// },
}

func rebaseRun(cmd *cobra.Command, args []string) {

	rebase := func(dir string, submodules Modules) {
    	for i := 0; i < len(submodules); i++ {

			status, err := submodules[i].Status()
			CheckIfError(err)

			if interactive {
				command, err := prompt.PromptForInteractive(args, submodules[i])
				CheckIfError(err)

				err = bash.RebaseViaBash(dir+"/"+*&status.Path, command)
				CheckIfError(err)

				continue
			}

			Info("Entering "+*&status.Path+"...")
			err = bash.Rebase(dir+"/"+*&status.Path, args)
			CheckIfError(err)
		}
    }

	for index, arg := range rebaseBoolArgIndexMap {
		rebaseBoolArgIndexMap[index] = arg.SetValue(rebaseBoolArgs[index])
	}
	for index, arg := range rebaseStringArgIndexMap {
		rebaseStringArgIndexMap[index] = arg.SetValue(rebaseStringArgs[index])
	}

	_args := Args.GenerateCommand(rebaseStringArgIndexMap, rebaseBoolArgIndexMap)
	args = append(_args, args...)

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir, dir)
	CheckIfError(err)

	if submodule != "" {
		
		_submodule, err := FindSubmodule(submodules, submodule)
		CheckIfError(err)

		status, err := _submodule.Status()
		CheckIfError(err)

		submodules, err = GetSubmodules(dir+"/"+*&status.Path, dir)
		rebase(dir, submodules)

		Info("Entering "+*&status.Path+"...")
		err = bash.Rebase(dir+"/"+*&status.Path, args)
		CheckIfError(err)

		return 
	}

	rebase(dir, submodules)

	Info("Entering /...")
	err = bash.Rebase(dir, args)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(rebaseCmd)
	
	rebaseStringArgIndexMap[0] = Args.StringArg{ Long: "signed", Short: "" } 
	rebaseStringArgIndexMap[1] = Args.StringArg{ Long: "strategy", Short: "s" } 
	rebaseStringArgIndexMap[2] = Args.StringArg{ Long: "strategy-option", Short: "X" } 
	rebaseStringArgIndexMap[3] = Args.StringArg{ Long: "C", Short: "C" } 
	rebaseStringArgIndexMap[4] = Args.StringArg{ Long: "whitespace", Short: "" } 
	rebaseStringArgIndexMap[5] = Args.StringArg{ Long: "rebase-merges", Short: "r" } 
	rebaseStringArgIndexMap[6] = Args.StringArg{ Long: "exec", Short: "x" } 
	rebaseStringArgIndexMap[7] = Args.StringArg{ Long: "onto", Short: "", NoEqual: true } 

	for index, val := range rebaseStringArgIndexMap {
		rebaseCmd.Flags().StringVarP(&rebaseStringArgs[index], val.Long, val.Short, "",  DocRoot+"/"+rebaseSlug+"#"+rebaseSlug+"-"+val.Long)
	}

	rebaseBoolArgIndexMap[0] = Args.BoolArg{ Long: "continue", Short: "" }
	rebaseBoolArgIndexMap[1] = Args.BoolArg{ Long: "abort", Short: "" }
	rebaseBoolArgIndexMap[2] = Args.BoolArg{ Long: "quit", Short: "" }
	rebaseBoolArgIndexMap[3] = Args.BoolArg{ Long: "keep-empty", Short: "" }
	rebaseBoolArgIndexMap[4] = Args.BoolArg{ Long: "allow-empty-message", Short: "" }
	rebaseBoolArgIndexMap[5] = Args.BoolArg{ Long: "skip", Short: "" }
	rebaseBoolArgIndexMap[6] = Args.BoolArg{ Long: "edit-todo", Short: "" }
	rebaseBoolArgIndexMap[7] = Args.BoolArg{ Long: "show-current-patch", Short: "" }
	rebaseBoolArgIndexMap[8] = Args.BoolArg{ Long: "merge", Short: "m" }
	rebaseBoolArgIndexMap[9] = Args.BoolArg{ Long: "quiet", Short: "q" } 
	rebaseBoolArgIndexMap[10] = Args.BoolArg{ Long: "verbose", Short: "v" } 
	rebaseBoolArgIndexMap[11] = Args.BoolArg{ Long: "stat", Short: "n" } 
	rebaseBoolArgIndexMap[12] = Args.BoolArg{ Long: "no-stat", Short: "" } 
	rebaseBoolArgIndexMap[13] = Args.BoolArg{ Long: "verify", Short: "" }
	rebaseBoolArgIndexMap[14] = Args.BoolArg{ Long: "no-verify", Short: "" }
	rebaseBoolArgIndexMap[15] = Args.BoolArg{ Long: "no-ff", Short: "" }
	rebaseBoolArgIndexMap[16] = Args.BoolArg{ Long: "force-rebase", Short: "f" }
	rebaseBoolArgIndexMap[17] = Args.BoolArg{ Long: "fork-point", Short: "" }
	rebaseBoolArgIndexMap[18] = Args.BoolArg{ Long: "no-fork-point", Short: "" }
	rebaseBoolArgIndexMap[19] = Args.BoolArg{ Long: "ignore-whitespace", Short: "" }
	rebaseBoolArgIndexMap[20] = Args.BoolArg{ Long: "committer-date-is-author-date", Short: "" }
	rebaseBoolArgIndexMap[21] = Args.BoolArg{ Long: "ignore-date", Short: "" }
	rebaseBoolArgIndexMap[22] = Args.BoolArg{ Long: "signoff", Short: "" }
	rebaseBoolArgIndexMap[23] = Args.BoolArg{ Long: "interactive", Short: "i" }
	rebaseBoolArgIndexMap[24] = Args.BoolArg{ Long: "preserve-merges", Short: "p" }
	rebaseBoolArgIndexMap[25] = Args.BoolArg{ Long: "root", Short: "" }
	rebaseBoolArgIndexMap[26] = Args.BoolArg{ Long: "autosquash", Short: "" }
	rebaseBoolArgIndexMap[27] = Args.BoolArg{ Long: "no-autosquash", Short: "" }

	for index, val := range rebaseBoolArgIndexMap {
		rebaseCmd.Flags().BoolVarP(&rebaseBoolArgs[index], val.Long, val.Short, false,  DocRoot+"/"+rebaseSlug+"#"+rebaseSlug+"-"+val.Long)
	}
}
