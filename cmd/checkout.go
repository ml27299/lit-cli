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
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	updateRun(cmd, append(args, []string{"silent", "prompt"}...))
	// },
}

func checkoutRun(cmd *cobra.Command, args []string) {

	checkout := func(dir string, submodules Modules) {
    	for i := 0; i < len(submodules); i++ {

    		if debug {println("Getting submodule tree")}
			status, err := submodules[i].Status()
			CheckIfError(err)

			if debug {println("Checking if interactive")}
			if interactive {

				if debug {println("Prompting interactive")}
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
    }

	for index, arg := range checkoutBoolArgIndexMap {
		arg.SetValue(checkoutBoolArgs[index])
	}
	for index, arg := range checkoutStringArgIndexMap {
		arg.SetValue(checkoutStringArgs[index])
	}

	if debug {println("Generating command")}
	_args := Args.GenerateCommand(checkoutStringArgIndexMap, checkoutBoolArgIndexMap)
	args = append(_args, args...)

	if debug {println("Finding root directory")}
	dir, err := paths.FindRootDir()
	CheckIfError(err)

	if debug {println("Getting submodules")}
	submodules, err := GetSubmodules(dir, dir)
	CheckIfError(err)

	if submodule != "" {
		
		if debug {println("Finding submodule : " + submodule)}
		_submodule, err := FindSubmodule(submodules, submodule)
		CheckIfError(err)

		if debug {println("Getting submodule tree")}
		status, err := _submodule.Status()
		CheckIfError(err)

		if debug {println("Getting submodules within "+submodule)}
		submodules, err = GetSubmodules(dir+"/"+*&status.Path, dir)
		checkout(dir, submodules)

		Info("Entering "+*&status.Path+"...")
		err = bash.Checkout(dir+"/"+*&status.Path, args)
		CheckIfError(err)

		return 
	}

	checkout(dir, submodules)

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
