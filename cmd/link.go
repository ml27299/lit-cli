package cmd

import (
	"os"
	"io"
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

	s_fileinfo, s_err := os.Stat(link.Source)
	d_fileinfo, d_err := os.Stat(link.Dest)

	CheckIfError(s_err)

	if d_err == nil {
		if !os.SameFile(s_fileinfo, d_fileinfo) {

			os.Remove(link.Dest)
			os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)

			if debug {println("S:"+ link.Source)}
			if debug {println("D:"+ link.Dest)}
			os.Link(link.Source, link.Dest)

		}else {
			// os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)
			// os.Link(link.Source, link.Dest)
		}
	} else {

		os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)
		os.Link(link.Source, link.Dest)
	}

}

func linkRun(cmd *cobra.Command, args []string) {

	dir, err := paths.FindRootDir()
	CheckIfError(err)

	config_files, err := paths.FindConfig(dir)
	CheckIfError(err)

	for _, config_file := range config_files {

		config_file_dir := filepath.Dir(config_file)
		err = os.Chdir(config_file_dir)
		CheckIfError(err)

		info, err := parser.ConfigViaPath(config_file_dir, dir)
		CheckIfError(err)

		links, err := info.GetLinks(dir)
		CheckIfError(err)
		
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
