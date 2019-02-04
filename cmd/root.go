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
	"os"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"../helpers/parser"
	//"../helpers/parser/paths"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "project",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.project.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func AddMissingFiles() error {

// 	dir, err := paths.FindRootDir()
// 	if err != nil{
// 		return err
// 	} 

// 	sourceMap, err := parser.MapSource(dir+"/lit.config.json")
// 	destMap, err := parser.MapDest(dir+"/lit.config.json")
// 	submodules, err := GetSubmodules(dir)

// 	if err != nil{
// 		return err
// 	} 

// 	var uneven_keys []string
// 	for key, val := range sourceMap {
// 		if len(sourceMap[key]) < len(destMap[key]) {
// 			uneven_keys = append(uneven_keys, key)
// 		} 
// 	}

// 	if len(uneven_keys) == 0 {
// 		return nil
// 	}

// 	links, err := parser.ParseConfig(dir+"/lit.config.json")
// 	if err != nil{
// 		return err
// 	} 

// 	type Missing struct {
// 		Value string
// 		Submodule_Paths []string
// 	}

// 	var missing_files []Missing
// 	for _, uneven_key := range uneven_keys {
// 		for  _, dests := range destMap[uneven_key] {

// 			for  _, dest := range dests {
// 				var found bool
// 				for  _, link := range links {
// 					if link.OG_Dest == dest {
// 						found = true
// 						break
// 					}
// 				}

// 				if !found {
// 					paths, err := FindAssociateGitModulePaths(submodules, sourceMap[uneven_key])
// 					missing_files = append(missing_files, Missing{
// 						Value: dest,
// 						Submodule_Paths: paths,
// 					})
// 				}
// 			}
// 		}
// 	}

// 	for _, missing_file := range missing_files {
// 		fmt.Println(missing_file.Value)
// 	}

// 	return nil
// }

// func AddMissingFiles() error {
// 	dir, err := paths.FindRootDir()
// 	if err != nil{
// 		return err
// 	} 

// 	items, err  := source.ParseConfig(dir+"/lit.config.json")
// 	if err != nil {
// 		return response, err
// 	}	

// 	paths,err := paths.Find()


// 	sourceItems, err := FindSourceItems()
// 	destItems, err := FindDestItems()
// 	submodules, err := GetSubmodules(dir)

// 	if err != nil {
// 		return err
// 	}

// 	var items []parser.Item
// 	if len(sourceItems) > len(destItems) {
// 		items = sourceItems
// 		fmt.Println("sourceItems")
// 	}else if(len(sourceItems) < len(destItems)){	
// 		items = destItems
// 		fmt.Println("destItems")
// 	}else {
// 		return nil
// 	}

// 	for _, item := range items {
// 		fmt.Println("S:"+item.Source)
// 		fmt.Println("D:"+item.Dest)
// 		modules, err := FindAssociateGitModules(item.Source, submodules)
// 		if err != nil {
// 			return err
// 		}

// 		fmt.Println(len(modules))

// 		if len(modules) == 1 {

// 			Copy(item.Dest, item.Source)
// 			os.Remove(item.Dest)
// 			Link(item)

// 		}else if len(modules) == 0 {
// 			continue
// 		}else if len(modules) > 1 {

// 			rl, err := readline.New("found more than 1 modules")
// 			if err != nil {
// 				panic(err)
// 			}
// 			defer rl.Close()

// 			for {
// 				line, err := rl.Readline()
// 				if err != nil { // io.EOF
// 					break
// 				}
// 				println(line)
// 			}
// 		}
// 	}

// 	return nil
// }

func FindHardLinkedFilePaths() ([]string, error){
	var response []string
	info, err  := parser.Config()
	if err != nil {
		return response, err
	}	

	links, err := info.GetLinks()
	for _,link := range links {
		response = append(response, link.Dest)
	}

	return response, nil
}

// func FindDestItems() ([]parser.Item, error){
// 	var response []parser.Item
// 	dir, err := paths.FindRootDir()
// 	if err != nil{
// 		return response, err
// 	} 

// 	items, err  := dest.ParseConfig(dir+"/lit.config.json")
// 	if err != nil {
// 		return response, err
// 	}	

// 	return items, nil
// }

// func FindSourceItems() ([]parser.Item, error){
// 	var response []parser.Item
// 	dir, err := paths.FindRootDir()
// 	if err != nil{
// 		return response, err
// 	} 

// 	items, err  := source.ParseConfig(dir+"/lit.config.json")
// 	if err != nil {
// 		return response, err
// 	}	

// 	return items, nil
// }

// func FindHardLinkedFilePaths() ([]string, error){
// 	var response []string
// 	dir, err := paths.FindRootDir()
// 	if err != nil{
// 		return response, err
// 	} 

// 	items, err  := source.ParseConfig(dir+"/lit.config.json")
// 	if err != nil {
// 		return response, err
// 	}	

// 	for _,item := range items {
// 		response = append(response, item.Dest)
// 	}

// 	return response, nil
// }

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
