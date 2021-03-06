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
	pullSlug = "git-pull"
	pullStringArgs [13]string
	pullBoolArgs [36]bool
	pullStringArgIndexMap = make(map[int]Args.StringArg)
	pullBoolArgIndexMap = make(map[int]Args.BoolArg)
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: DocRoot+"/"+pullSlug,
	Long: `ex. lit pull origin master`,
	Run: pullRun,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	updateRun(cmd, append(args, []string{"silent", "prompt"}...))
	// },
	PostRun: func(cmd *cobra.Command, args []string) {
		linkRun(cmd, args)
	},
}


func pullRun(cmd *cobra.Command, args []string) {

	pull := func(dir string, submodules Modules) {
    	for i := 0; i < len(submodules); i++ {

			status, err := submodules[i].Status()
			CheckIfError(err)

			if interactive {
				command, err := prompt.PromptForInteractive(args, submodules[i])
				CheckIfError(err)

				err = bash.PullViaBash(dir+"/"+*&status.Path, command)
				CheckIfError(err)

				continue
			}

			Info("Entering "+*&status.Path+"...")
			err = bash.Pull(dir+"/"+*&status.Path, args)
			CheckIfError(err)
		}
    }

	for index, arg := range pullBoolArgIndexMap {
		pullBoolArgIndexMap[index] = arg.SetValue(pullBoolArgs[index])
	}
	for index, arg := range pullStringArgIndexMap {
		pullStringArgIndexMap[index] = arg.SetValue(pullStringArgs[index])
	}

	_args := Args.GenerateCommand(pullStringArgIndexMap, pullBoolArgIndexMap)
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
		pull(dir, submodules)

		Info("Entering "+*&status.Path+"...")
		err = bash.Pull(dir+"/"+*&status.Path, args)
		CheckIfError(err)

		return 
	}

	pull(dir, submodules)

	Info("Entering /...")
	err = bash.Pull(dir+"/", args)
	CheckIfError(err)

	SyncCommitIds(submodules, dir)

	// for i := 0; i < len(submodules); i++ {
	// 	status, err := submodules[i].Status()
	// 	CheckIfError(err)

	// 	bash.AddViaBash(dir, status.Path)
	// }

	// changes, err := bash.HasCommitChanges(dir)
	// CheckIfError(err)
	// if changes {
	// 	bash.CommitViaBash(dir, "-m \"synced commit id\"")
	// }
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullStringArgIndexMap[0] = Args.StringArg{ Long: "recurse-submodules", Short: "" } 
	pullStringArgIndexMap[1] = Args.StringArg{ Long: "no-recurse-submodules", Short: "" } 
	pullStringArgIndexMap[2] = Args.StringArg{ Long: "gpg-sign", Short: "S" } 
	pullStringArgIndexMap[3] = Args.StringArg{ Long: "log", Short: "" } 
	pullStringArgIndexMap[4] = Args.StringArg{ Long: "strategy", Short: "s" } 
	pullStringArgIndexMap[5] = Args.StringArg{ Long: "strategy-option", Short: "X" } 
	pullStringArgIndexMap[6] = Args.StringArg{ Long: "depth", Short: "" } 
	pullStringArgIndexMap[7] = Args.StringArg{ Long: "deepen", Short: "" } 
	pullStringArgIndexMap[8] = Args.StringArg{ Long: "shallow-since", Short: "" } 
	pullStringArgIndexMap[9] = Args.StringArg{ Long: "shallow-exclude", Short: "" } 
	pullStringArgIndexMap[10] = Args.StringArg{ Long: "negotiation-tip", Short: "" } 
	pullStringArgIndexMap[11] = Args.StringArg{ Long: "server-option", Short: "o" } 
	pullStringArgIndexMap[12] = Args.StringArg{ Long: "upload-pack", Short: "", NoEqual: true } 
	
	for index, val := range pullStringArgIndexMap {
		pullCmd.Flags().StringVarP(&pullStringArgs[index], val.Long, val.Short, "",  DocRoot+"/"+pullSlug+"#"+pullSlug+"-"+val.Long)
	}

	pullBoolArgIndexMap[0] = Args.BoolArg{ Long: "quiet", Short: "q" } 
	pullBoolArgIndexMap[1] = Args.BoolArg{ Long: "verbose", Short: "v" } 
	pullBoolArgIndexMap[2] = Args.BoolArg{ Long: "commit", Short: "" } 
	pullBoolArgIndexMap[3] = Args.BoolArg{ Long: "no-commit", Short: "" } 
	pullBoolArgIndexMap[4] = Args.BoolArg{ Long: "edit", Short: "e" } 
	pullBoolArgIndexMap[5] = Args.BoolArg{ Long: "no-edit", Short: "" } 
	pullBoolArgIndexMap[6] = Args.BoolArg{ Long: "ff", Short: "" } 
	pullBoolArgIndexMap[7] = Args.BoolArg{ Long: "no-ff", Short: "" } 
	pullBoolArgIndexMap[8] = Args.BoolArg{ Long: "ff-only", Short: "" } 
	pullBoolArgIndexMap[9] = Args.BoolArg{ Long: "no-log", Short: "" } 
	pullBoolArgIndexMap[10] = Args.BoolArg{ Long: "signoff", Short: "" } 
	pullBoolArgIndexMap[11] = Args.BoolArg{ Long: "no-signoff", Short: "" } 
	pullBoolArgIndexMap[12] = Args.BoolArg{ Long: "stat", Short: "n" } 
	pullBoolArgIndexMap[13] = Args.BoolArg{ Long: "no-stat", Short: "" } 
	pullBoolArgIndexMap[14] = Args.BoolArg{ Long: "squash", Short: "" } 
	pullBoolArgIndexMap[15] = Args.BoolArg{ Long: "no-squash", Short: "" } 
	pullBoolArgIndexMap[16] = Args.BoolArg{ Long: "verify-signatures", Short: "" } 
	pullBoolArgIndexMap[17] = Args.BoolArg{ Long: "no-verify-signatures", Short: "" } 
	pullBoolArgIndexMap[18] = Args.BoolArg{ Long: "summary", Short: "" } 
	pullBoolArgIndexMap[19] = Args.BoolArg{ Long: "no-summary", Short: "" } 
	pullBoolArgIndexMap[20] = Args.BoolArg{ Long: "allow-unrelated-histories", Short: "" } 
	pullBoolArgIndexMap[21] = Args.BoolArg{ Long: "rebase", Short: "r" } 
	pullBoolArgIndexMap[22] = Args.BoolArg{ Long: "no-rebase", Short: "" } 
	pullBoolArgIndexMap[23] = Args.BoolArg{ Long: "autostash", Short: "" } 
	pullBoolArgIndexMap[24] = Args.BoolArg{ Long: "no-autostash", Short: "" }
	pullBoolArgIndexMap[25] = Args.BoolArg{ Long: "all", Short: "" }
	pullBoolArgIndexMap[26] = Args.BoolArg{ Long: "append", Short: "a" }
	pullBoolArgIndexMap[27] = Args.BoolArg{ Long: "unshallow", Short: "" }
	pullBoolArgIndexMap[28] = Args.BoolArg{ Long: "update-shallow", Short: "" }
	pullBoolArgIndexMap[29] = Args.BoolArg{ Long: "force", Short: "f" }
	pullBoolArgIndexMap[30] = Args.BoolArg{ Long: "keep", Short: "k" }
	pullBoolArgIndexMap[31] = Args.BoolArg{ Long: "no-tags", Short: "" }
	pullBoolArgIndexMap[32] = Args.BoolArg{ Long: "update-head-ok", Short: "u" }
	pullBoolArgIndexMap[33] = Args.BoolArg{ Long: "progress", Short: "" }
	pullBoolArgIndexMap[34] = Args.BoolArg{ Long: "ipv4", Short: "4" }
	pullBoolArgIndexMap[35] = Args.BoolArg{ Long: "ipv6", Short: "6" }

	for index, val := range pullBoolArgIndexMap {
		pullCmd.Flags().BoolVarP(&pullBoolArgs[index], val.Long, val.Short, false,  DocRoot+"/"+pullSlug+"#"+pullSlug+"-"+val.Long)
	}
}
