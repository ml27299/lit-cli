package cmd

import (
	"github.com/spf13/cobra"
	. "../helpers"
	"../helpers/paths"
	"../helpers/bash"
	Args "../helpers/args"
	"../helpers/prompt"
)

var (
	checkoutSlug = "git-checkout"
	checkoutStringArgs [2]string
	checkoutBoolArgs [18]bool
	checkoutStringArgIndexMap = make(map[int]Args.StringArg)
	checkoutBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short:  DocRoot+"/"+checkoutSlug,
	Long: `ex. lit checkout master`,
	Run: checkoutRun,
}

func checkoutRun(cmd *cobra.Command, args []string) {

	for index, arg := range checkoutBoolArgIndexMap {
		arg.SetValue(checkoutBoolArgs[index])
	}
	for index, arg := range checkoutStringArgIndexMap {
		arg.SetValue(checkoutStringArgs[index])
	}

	_args := Args.GenerateCommand(checkoutStringArgIndexMap, checkoutBoolArgIndexMap)
	args = append(_args, args...)

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	if submodule != "" {
		
		_submodule, err := FindSubmodule(submodules, submodule)
		CheckIfError(err)

		status, err := _submodule.Status()
		CheckIfError(err)

		err = bash.Checkout(dir+"/"+*&status.Path, args)
		CheckIfError(err)

		return 
	}

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)

		if interactive {
			command, err := prompt.PromptForInteractive(args, submodules[i])
			CheckIfError(err)

			err = bash.CheckoutViaBash(dir+"/"+*&status.Path, command)
			CheckIfError(err)

			continue
		}

		Info("Entering "+*&status.Path+"...")
		err = bash.Checkout(dir+"/"+*&status.Path, args)
		CheckIfError(err)
	}

	Info("Entering /...")
	err = bash.Checkout(dir, args)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(checkoutCmd)

	checkoutStringArgIndexMap[0] = Args.StringArg{ Long: "conflict", Short: "" } 
	checkoutStringArgIndexMap[1] = Args.StringArg{ Long: "orphan", Short: "", NoEqual: true } 

	for index, val := range checkoutStringArgIndexMap {
		checkoutCmd.Flags().StringVarP(&checkoutStringArgs[index], val.Long, val.Short, "",  DocRoot+"/"+checkoutSlug+"#"+checkoutSlug+"-"+val.Long)
	}

	checkoutBoolArgIndexMap[0] = Args.BoolArg{ Long: "quiet", Short: "q" } 
	checkoutBoolArgIndexMap[1] = Args.BoolArg{ Long: "progress", Short: "" } 
	checkoutBoolArgIndexMap[2] = Args.BoolArg{ Long: "no-progress", Short: "" } 
	checkoutBoolArgIndexMap[3] = Args.BoolArg{ Long: "force", Short: "f" }
	checkoutBoolArgIndexMap[4] = Args.BoolArg{ Long: "ours", Short: "" }
	checkoutBoolArgIndexMap[5] = Args.BoolArg{ Long: "theirs", Short: "" }
	checkoutBoolArgIndexMap[6] = Args.BoolArg{ Long: "b", Short: "b" }
	checkoutBoolArgIndexMap[7] = Args.BoolArg{ Long: "B", Short: "B" }
	checkoutBoolArgIndexMap[8] = Args.BoolArg{ Long: "track", Short: "t" }
	checkoutBoolArgIndexMap[9] = Args.BoolArg{ Long: "no-track", Short: "" }
	checkoutBoolArgIndexMap[10] = Args.BoolArg{ Long: "l", Short: "l" }
	checkoutBoolArgIndexMap[11] = Args.BoolArg{ Long: "detach", Short: "" }
	checkoutBoolArgIndexMap[12] = Args.BoolArg{ Long: "ignore-skip-worktree-bits", Short: "" }
	checkoutBoolArgIndexMap[13] = Args.BoolArg{ Long: "merge", Short: "m" }
	checkoutBoolArgIndexMap[14] = Args.BoolArg{ Long: "patch", Short: "p" }
	checkoutBoolArgIndexMap[15] = Args.BoolArg{ Long: "ignore-other-worktrees", Short: "" }
	checkoutBoolArgIndexMap[16] = Args.BoolArg{ Long: "recurse-submodules", Short: "" }
	checkoutBoolArgIndexMap[17] = Args.BoolArg{ Long: "no-recurse-submodules", Short: "" }

	for index, val := range checkoutBoolArgIndexMap {
		checkoutCmd.Flags().BoolVarP(&checkoutBoolArgs[index], val.Long, val.Short, false,  DocRoot+"/"+checkoutSlug+"#"+checkoutSlug+"-"+val.Long)
	}
}
