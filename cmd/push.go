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
	pushSlug = "git-push"
	pushStringArgs [7]string
	pushBoolArgs [23]bool
	pushStringArgIndexMap = make(map[int]Args.StringArg)
	pushBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: DocRoot+"/"+pushSlug,
	Long: `ex. lit push origin master`,
	Run: pushRun,
}

func pushRun(cmd *cobra.Command, args []string) {

	for index, arg := range pushBoolArgIndexMap {
		pushBoolArgIndexMap[index] = arg.SetValue(pushBoolArgs[index])
	}
	for index, arg := range pushStringArgIndexMap {
		pushStringArgIndexMap[index] = arg.SetValue(pushStringArgs[index])
	}

	_args := Args.GenerateCommand(pushStringArgIndexMap, pushBoolArgIndexMap)
	args = append(_args, args...)

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)	

		if submodule != "" && status.Path == submodule {
			err = bash.Push(dir+"/"+*&status.Path, args)
			CheckIfError(err)

			break
		}

		if interactive {
			command, err := prompt.PromptForInteractive(args, submodules[i])
			CheckIfError(err)

			err = bash.PushViaBash(dir+"/"+*&status.Path, command)
			CheckIfError(err)

			continue
		}

		Info("Entering "+*&status.Path+"...")
		err = bash.Push(dir+"/"+*&status.Path, args)
		CheckIfError(err)
	}

	if submodule == "" {
		Info("Entering /...")
		err = bash.Push(dir, args)
		CheckIfError(err)
	}
}

func init() {
	rootCmd.AddCommand(pushCmd)
	
	pushStringArgIndexMap[0] = Args.StringArg{ Long: "signed", Short: "" } 
	pushStringArgIndexMap[1] = Args.StringArg{ Long: "push-option", Short: "o" } 
	pushStringArgIndexMap[2] = Args.StringArg{ Long: "receive-pack", Short: "" } 
	pushStringArgIndexMap[3] = Args.StringArg{ Long: "exec", Short: "" } 
	pushStringArgIndexMap[4] = Args.StringArg{ Long: "force-with-lease", Short: "" } 
	pushStringArgIndexMap[5] = Args.StringArg{ Long: "repo", Short: "" } 
	pushStringArgIndexMap[6] = Args.StringArg{ Long: "recurse-submodules", Short: "" } 

	for index, val := range pushStringArgIndexMap {
		pushCmd.Flags().StringVarP(&pushStringArgs[index], val.Long, val.Short, "",  DocRoot+"/"+pushSlug+"#"+pushSlug+"-"+val.Long)
	}

	pushBoolArgIndexMap[0] = Args.BoolArg{ Long: "all", Short: "" }
	pushBoolArgIndexMap[1] = Args.BoolArg{ Long: "prune", Short: "" }
	pushBoolArgIndexMap[2] = Args.BoolArg{ Long: "mirror", Short: "" }
	pushBoolArgIndexMap[3] = Args.BoolArg{ Long: "dry-run", Short: "n" }
	pushBoolArgIndexMap[4] = Args.BoolArg{ Long: "porcelain", Short: "" }
	pushBoolArgIndexMap[5] = Args.BoolArg{ Long: "delete", Short: "d" }
	pushBoolArgIndexMap[6] = Args.BoolArg{ Long: "tags", Short: "" }
	pushBoolArgIndexMap[7] = Args.BoolArg{ Long: "follow-tags", Short: "" }
	pushBoolArgIndexMap[8] = Args.BoolArg{ Long: "no-signed", Short: "" }
	pushBoolArgIndexMap[9] = Args.BoolArg{ Long: "atomic", Short: "" }
	pushBoolArgIndexMap[10] = Args.BoolArg{ Long: "no-atomic", Short: "" }
	pushBoolArgIndexMap[11] = Args.BoolArg{ Long: "no-force-with-lease", Short: "" }
	pushBoolArgIndexMap[12] = Args.BoolArg{ Long: "force", Short: "f" }
	pushBoolArgIndexMap[13] = Args.BoolArg{ Long: "set-upstream", Short: "u" }
	pushBoolArgIndexMap[14] = Args.BoolArg{ Long: "thin", Short: "" }
	pushBoolArgIndexMap[15] = Args.BoolArg{ Long: "no-thin", Short: "" }
	pushBoolArgIndexMap[16] = Args.BoolArg{ Long: "quiet", Short: "q" }
	pushBoolArgIndexMap[17] = Args.BoolArg{ Long: "verbose", Short: "v" }
	pushBoolArgIndexMap[18] = Args.BoolArg{ Long: "progress", Short: "" }
	pushBoolArgIndexMap[19] = Args.BoolArg{ Long: "verify", Short: "" }
	pushBoolArgIndexMap[20] = Args.BoolArg{ Long: "no-verify", Short: "" }
	pushBoolArgIndexMap[21] = Args.BoolArg{ Long: "ipv4", Short: "4" }
	pushBoolArgIndexMap[22] = Args.BoolArg{ Long: "ipv6", Short: "6" }

	for index, val := range pushBoolArgIndexMap {
		pushCmd.Flags().BoolVarP(&pushBoolArgs[index], val.Long, val.Short, false,  DocRoot+"/"+pushSlug+"#"+pushSlug+"-"+val.Long)
	}
}
