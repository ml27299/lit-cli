package cmd

import (
	"os"
	//"fmt"
	"io"
	"strings"
	"github.com/spf13/cobra"
	"github.com/ml27299/lit-cli/helpers/parser"
	"github.com/ml27299/lit-cli/helpers/paths"
	. "github.com/ml27299/lit-cli/helpers"
	"path/filepath"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "hard links files from a directory, requires a lit.config.json",
	Long: `ex. lit link`,
	Run: linkRun,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	updateRun(cmd, append(args, []string{"silent", "prompt"}...))
	// },
}

func Copy(src string, dst string) (int64, error) {
    source, err := os.Open(src)
    if err != nil {
        return 0, err
    }
    defer source.Close()
    
    destination, err := os.Create(dst)
    if err != nil {
        return 0, err
    }
    defer destination.Close()

    nBytes, err := io.Copy(destination, source)
    return nBytes, err
}

func Link(link parser.Link) {

	if debug {println("Doing link")}

	s_fileinfo, s_err := os.Stat(link.Source)
	d_fileinfo, d_err := os.Stat(link.Dest)

	CheckIfError(s_err)

	if debug {println("S:"+ link.Source)}
	if debug {println("D:"+ link.Dest)}

	if d_err == nil {
		if debug {println("d_err = nil")}
		if !os.SameFile(s_fileinfo, d_fileinfo) {
			if debug {println("not smae file")}
			os.Remove(link.Dest)
			os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)

			os.Link(link.Source, link.Dest)

		}else {
			if debug {println("smae file")}
			// os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)
			// os.Link(link.Source, link.Dest)
		}
	} else {
		if debug {println("d_err != nil")}
		os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)
		os.Link(link.Source, link.Dest)
	}

}


func linkRun(cmd *cobra.Command, args []string) {
	if debug {println("Starting link")}

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	if debug {println("root dir:"+dir)}

	config_files, err := paths.FindConfig(dir)
	CheckIfError(err)

	link_map := make(map[string][]parser.Link)

	for _, config_file := range config_files {

		if debug {println("Config file:"+config_file)}

		config_file_dir := filepath.Dir(config_file)
		err = os.Chdir(config_file_dir)
		CheckIfError(err)

		if debug {println("Config file dir:"+config_file_dir)}

		info, err := parser.ConfigViaPath(config_file_dir, dir)
		CheckIfError(err)

		links, err := info.GetLinks(dir)
		CheckIfError(err)
		if debug {println(links)}

		link_map[config_file_dir] = links

	}

	for config_file_dir, links := range link_map {
		for config_file_dir2, _ := range link_map {
			if config_file_dir2 != config_file_dir {
				for _, link := range links {
					var links_to_add []parser.Link
					if strings.Contains(link.Dest, config_file_dir2) == true {
						links_to_add = append(links_to_add, link)
					}
					if len(links_to_add) > 0 {
						link_map[config_file_dir2] = parser.AppendUniqueLinks(link_map[config_file_dir2], links_to_add)
					}
				}
			}
		}
	}

	for config_file_dir, links := range link_map {
		err = UpdateGitignore(config_file_dir, links)
		CheckIfError(err)

		for _, link := range links {
			Link(link)
		}
	}
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
