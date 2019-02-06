package cmd

import (
	"os"
	"io"
	"github.com/spf13/cobra"
	"../helpers/parser"
	. "../helpers"
	"path/filepath"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "hard links files from a directory, requires a lit.config.json",
	Long: `ex. lit link`,
	Run: linkRun,
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

func Link(link parser.Link) error {

	s_fileinfo, s_err := os.Stat(link.Source)
	d_fileinfo, d_err := os.Stat(link.Dest)

	if s_err != nil {
		return s_err
	}

	if d_err == nil {
		if !os.SameFile(s_fileinfo, d_fileinfo) {

			os.Remove(link.Dest)
			os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)
			os.Link(link.Source, link.Dest)

		}else {
			os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)
			os.Link(link.Source, link.Dest)
		}
	} else {

		os.MkdirAll(filepath.Dir(link.Dest), os.ModePerm)
		os.Link(link.Source, link.Dest)
	}

	return nil
}

func linkRun(cmd *cobra.Command, args []string) {

	info, err  := parser.Config()
	CheckIfError(err)

	links, err := info.GetLinks()
	CheckIfError(err)

	for _, link := range links {
		err := Link(link)
		CheckIfError(err)
	}
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
