// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"
	"os"
	"github.com/ml27299/lit-cli/helpers"
	"github.com/ml27299/lit-cli/helpers/bash"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	DocRoot = "https://git-scm.com/docs"
	cfgFile string
	interactive bool
	submodule string
	debug bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lit",
	Short: "The lit cli can be used to build modular applications utilizing an architecture built around hard links and git submodules",
	Long: ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.project.yaml)")
	rootCmd.PersistentFlags().BoolVar(&interactive, "inter", false, "Run lit in interactivce mode")
	rootCmd.PersistentFlags().StringVar(&submodule, "submodule", "", "run a lit command for one submodule")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Get a verbose log of whats lit is doing")
}


func SyncCommitIds(submodules helpers.Modules, dir string) {
	
	submoduleMap := make(map[string]string)

	for i := 0; i < len(submodules); i++ {
		
		status, err := submodules[i].Status()
		helpers.CheckIfError(err)

		for z := 0; z < len(submodules); z++ {
			if submodules[z] != submodules[i] {

				status2, err := submodules[z].Status()
				helpers.CheckIfError(err)

				if len(status.Path) > len(status2.Path) {
					if strings.Contains(status.Path, status2.Path) == true {
						submoduleMap[status.Path] = status2.Path
					}
				}else {
					if strings.Contains(status2.Path, status.Path) == true {
						submoduleMap[status2.Path] =  status.Path
					}
				}
			}
		}
	}

	for i := 0; i < len(submodules); i++ {

		status, err := submodules[i].Status()
		helpers.CheckIfError(err)

		if len(submoduleMap[status.Path]) > 0 {
			bash.AddViaBash(dir+"/"+submoduleMap[status.Path], strings.Replace(status.Path, submoduleMap[status.Path]+"/", "", 1))
		}else {
			bash.AddViaBash(dir, status.Path)
		}
	}

	for _, modulePath := range submoduleMap {
		changes, err := bash.HasCommitChanges(dir+"/"+modulePath)
		helpers.CheckIfError(err)
		if changes {
			helpers.Info("Entering /"+modulePath)
			bash.CommitViaBash(dir+"/"+modulePath, "-m \"synced commit id\"")
		}
	}

	for h := 0; h < len(submodules); h++ {

		status, err := submodules[h].Status()
		helpers.CheckIfError(err)

		if len(submoduleMap[status.Path]) == 0 {
			bash.AddViaBash(dir, status.Path)
		}
	}

	changes, err := bash.HasCommitChanges(dir)
	helpers.CheckIfError(err)
	if changes {
		helpers.Info("Entering /...")
		bash.CommitViaBash(dir, "-m \"synced commit id\"")
	}

}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".project" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".project")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
