package cmd

import (
	"github.com/spf13/cobra"
	. "../helpers"
	"../helpers/paths"
	"../helpers/bash"
	Args "../helpers/args"
)

var (
	mergeSlug = "git-merge"
	mergeStringArgs [6]string
	mergeBoolArgs [27]bool
	mergeStringArgIndexMap = make(map[int]Args.StringArg)
	mergeBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: DocRoot+"/"+mergeSlug,
	Long: `ex. lit merge master`,
	Run: mergeRun,
	PostRun: linkRun,
}

func mergeRun(cmd *cobra.Command, args []string) {

	for index, arg := range mergeBoolArgIndexMap {
		arg.SetValue(mergeBoolArgs[index])
	}
	for index, arg := range mergeStringArgIndexMap {
		arg.SetValue(mergeStringArgs[index])
	}	

	_args := Args.GenerateCommand(mergeStringArgIndexMap, mergeBoolArgIndexMap)
	args = append(_args, args...)

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	submodules, err := GetSubmodules(dir)
	CheckIfError(err)

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		CheckIfError(err)

		Info("Entering "+*&status.Path+"...")

		err = bash.Merge(dir+"/"+*&status.Path, args)
    	CheckIfError(err)
	}

	Info("Entering /...")

	err = bash.Merge(dir+"/", args)
	CheckIfError(err)
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	mergeStringArgIndexMap[0] = Args.StringArg{ Long: "gpg-sign", Short: "S" } 
	mergeStringArgIndexMap[1] = Args.StringArg{ Long: "log", Short: "" } 
	mergeStringArgIndexMap[2] = Args.StringArg{ Long: "strategy", Short: "s" } 
	mergeStringArgIndexMap[3] = Args.StringArg{ Long: "strategy-option", Short: "X" } 
	mergeStringArgIndexMap[4] = Args.StringArg{ Long: "m", Short: "m" } 
	mergeStringArgIndexMap[5] = Args.StringArg{ Long: "file", Short: "F" } 
	
	for index, val := range mergeStringArgIndexMap {
		mergeCmd.Flags().StringVarP(&mergeStringArgs[index], val.Long, val.Short, "",  DocRoot+"/"+mergeSlug+"#"+mergeSlug+"-"+val.Long)
	}

	mergeBoolArgIndexMap[0] = Args.BoolArg{ Long: "commit", Short: "" } 
	mergeBoolArgIndexMap[1] = Args.BoolArg{ Long: "no-commit", Short: "" } 
	mergeBoolArgIndexMap[2] = Args.BoolArg{ Long: "edit", Short: "e" } 
	mergeBoolArgIndexMap[3] = Args.BoolArg{ Long: "no-edit", Short: "" } 
	mergeBoolArgIndexMap[4] = Args.BoolArg{ Long: "ff", Short: "" } 
	mergeBoolArgIndexMap[5] = Args.BoolArg{ Long: "no-ff", Short: "" } 
	mergeBoolArgIndexMap[6] = Args.BoolArg{ Long: "ff-only", Short: "" } 
	mergeBoolArgIndexMap[7] = Args.BoolArg{ Long: "no-log", Short: "" } 
	mergeBoolArgIndexMap[8] = Args.BoolArg{ Long: "signoff", Short: "" } 
	mergeBoolArgIndexMap[9] = Args.BoolArg{ Long: "no-signoff", Short: "" } 
	mergeBoolArgIndexMap[10] = Args.BoolArg{ Long: "stat", Short: "n" } 
	mergeBoolArgIndexMap[11] = Args.BoolArg{ Long: "no-stat", Short: "" } 
	mergeBoolArgIndexMap[12] = Args.BoolArg{ Long: "squash", Short: "" } 
	mergeBoolArgIndexMap[13] = Args.BoolArg{ Long: "no-squash", Short: "" } 
	mergeBoolArgIndexMap[14] = Args.BoolArg{ Long: "verify-signatures", Short: "" } 
	mergeBoolArgIndexMap[15] = Args.BoolArg{ Long: "no-verify-signatures", Short: "" } 
	mergeBoolArgIndexMap[16] = Args.BoolArg{ Long: "summary", Short: "" } 
	mergeBoolArgIndexMap[17] = Args.BoolArg{ Long: "no-summary", Short: "" } 
	mergeBoolArgIndexMap[18] = Args.BoolArg{ Long: "quiet", Short: "q" } 
	mergeBoolArgIndexMap[19] = Args.BoolArg{ Long: "verbose", Short: "v" } 
	mergeBoolArgIndexMap[20] = Args.BoolArg{ Long: "progress", Short: "" } 
	mergeBoolArgIndexMap[21] = Args.BoolArg{ Long: "no-progress", Short: "" } 
	mergeBoolArgIndexMap[22] = Args.BoolArg{ Long: "allow-unrelated-histories", Short: "" } 
	mergeBoolArgIndexMap[23] = Args.BoolArg{ Long: "rerere-autoupdate", Short: "" } 
	mergeBoolArgIndexMap[24] = Args.BoolArg{ Long: "no-rerere-autoupdate", Short: "" } 
	mergeBoolArgIndexMap[25] = Args.BoolArg{ Long: "abort", Short: "" } 
	mergeBoolArgIndexMap[26] = Args.BoolArg{ Long: "continue", Short: "" } 

	for index, val := range mergeBoolArgIndexMap {
		mergeCmd.Flags().BoolVarP(&mergeBoolArgs[index], val.Long, val.Short, false,  DocRoot+"/"+mergeSlug+"#"+mergeSlug+"-"+val.Long)
	}
}
