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
	statusSlug = "git-status"
	statusStringArgs [5]string
	statusBoolArgs [12]bool
	statusStringArgIndexMap = make(map[int]Args.StringArg)
	statusBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short:  DocRoot+"/"+statusSlug,
	Long: `ex. lit status`,
	Run: statusRun,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	updateRun(cmd, append(args, []string{"silent", "prompt"}...))
	// },
}

func statusRun(cmd *cobra.Command, args []string) {

	_status := func(dir string, submodules Modules) {
    	for i := 0; i < len(submodules); i++ {

			status, err := submodules[i].Status()
			CheckIfError(err)

			if interactive {
				command, err := prompt.PromptForInteractive(args, submodules[i])
				CheckIfError(err)

				err = bash.StatusViaBash(dir+"/"+*&status.Path, command)
				CheckIfError(err)

				continue
			}

			Info("Entering "+*&status.Path+"...")
			err = bash.Status(dir+"/"+*&status.Path, args)
			CheckIfError(err)
		}
    }

	for index, arg := range statusBoolArgIndexMap {
		statusBoolArgIndexMap[index] = arg.SetValue(statusBoolArgs[index])
	}
	for index, arg := range statusStringArgIndexMap {
		statusStringArgIndexMap[index] = arg.SetValue(statusStringArgs[index])
	}

	_args := Args.GenerateCommand(statusStringArgIndexMap, statusBoolArgIndexMap)
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
		_status(dir, submodules)

		Info("Entering "+*&status.Path+"...")
		err = bash.Status(dir+"/"+*&status.Path, args)
		CheckIfError(err)

		return 
	}

	_status(dir, submodules)

	Info("Entering /...")
	err = bash.Status(dir, args)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(statusCmd)
	
	statusStringArgIndexMap[0] = Args.StringArg{ Long: "porcelain", Short: "" } 
	statusStringArgIndexMap[1] = Args.StringArg{ Long: "ignore-submodules", Short: "" } 
	statusStringArgIndexMap[2] = Args.StringArg{ Long: "ignored", Short: "" } 
	statusStringArgIndexMap[3] = Args.StringArg{ Long: "column", Short: "" } 
	statusStringArgIndexMap[4] = Args.StringArg{ Long: "find-renames", Short: "" } 

	for index, val := range statusStringArgIndexMap {
		statusCmd.Flags().StringVarP(&statusStringArgs[index], val.Long, val.Short, "",  DocRoot+"/"+statusSlug+"#"+statusSlug+"-"+val.Long)
	}

	statusBoolArgIndexMap[0] = Args.BoolArg{ Long: "short", Short: "s" }
	statusBoolArgIndexMap[1] = Args.BoolArg{ Long: "z", Short: "z" } 
	statusBoolArgIndexMap[2] = Args.BoolArg{ Long: "branch", Short: "b" }
	statusBoolArgIndexMap[3] = Args.BoolArg{ Long: "show-stash", Short: "" }
	statusBoolArgIndexMap[4] = Args.BoolArg{ Long: "long", Short: "" }
	statusBoolArgIndexMap[5] = Args.BoolArg{ Long: "verbose", Short: "" }
	statusBoolArgIndexMap[6] = Args.BoolArg{ Long: "untracked-files", Short: "" }
	statusBoolArgIndexMap[7] = Args.BoolArg{ Long: "no-column", Short: "" }
	statusBoolArgIndexMap[8] = Args.BoolArg{ Long: "ahead-behind", Short: "" }
	statusBoolArgIndexMap[9] = Args.BoolArg{ Long: "no-ahead-behind", Short: "" }
	statusBoolArgIndexMap[10] = Args.BoolArg{ Long: "renames", Short: "" } 
	statusBoolArgIndexMap[11] = Args.BoolArg{ Long: "no-renames", Short: "" } 

	for index, val := range statusBoolArgIndexMap {
		statusCmd.Flags().BoolVarP(&statusBoolArgs[index], val.Long, val.Short, false,  DocRoot+"/"+statusSlug+"#"+statusSlug+"-"+val.Long)
	}
}
