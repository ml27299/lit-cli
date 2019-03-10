package helpers

import (
	"strings"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"github.com/ml27299/lit-cli/helpers/resources"
	"github.com/ml27299/lit-cli/helpers/parser"
)

func UpdateGitignore(path string, links []parser.Link) error {

	if _, err := os.Stat(path+"/.gitignore"); os.IsNotExist(err) {

		data, err := resources.Asset(".gitignore")
		file, err := os.Create(path+"/.gitignore")

		defer file.Close()
		if err != nil {
			return err
		}

		_, err = file.Write(data)
		if err != nil {
			return err
		}
	}

	err := cleanGeneratedContent(path+"/.gitignore")
	if err != nil {
		return err
	}

	var link_dests []string
	for _, link := range links {
		link_dests = append(link_dests, strings.Replace(link.Dest, path, "", -1))
	}

	err = generateContent(path+"/.gitignore", link_dests)
	if err != nil {
		return err
	}

	return nil
}

func cleanGeneratedContent(path string) error {
	lines, err := getLinesFromFile(path)
	if err != nil {
		return err
	}

	var (
		skip = false
		fileContent = ""
	)

	for _, line := range lines {
		if line == "### BEGIN GENERATED CONTENT" {
			fileContent += line
			fileContent += "\n"
			skip = true
		}

		if line == "### END GENERATED CONTENT" {
			skip = false
		}

		if skip {
			continue
		}

		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func generateContent(path string, links []string) error {
	lines, err := getLinesFromFile(path)
	if err != nil {
		return err
	}

	var (
		insert = false
		fileContent = ""
	)

	for _, line := range lines {
		if line == "### BEGIN GENERATED CONTENT" {
			fileContent += line
			fileContent += "\n"
			insert = true
		}

		if insert {
			for _, link := range links {
				fileContent += link
				fileContent += "\n"
			}
			insert = false
			continue
		}

		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func getLinesFromFile(filePath string) ([]string, error) {
	
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}