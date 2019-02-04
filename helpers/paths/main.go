package paths

import (
	//"fmt"
	"os"
	"path/filepath"
	"errors"
)

var paths []string
func Normalize(str string) (string, error) {
	var response string
	dir, err := FindRootDir()

	if err != nil {
		return response, err
	}
	
	if str[:1] == "/" {
		response = dir+str
	}else {
		response =  dir+"/"+str
	}
	
	if response[len(response)-1:] == "/" {
		response = response[:len(response)-2]
	}

	return response, nil
}

func Find(path string) ([]string, error) {
	paths = nil

	err := filepath.Walk(path, visitFind)
    if err != nil {
        return paths, err
    }

   	return paths, nil
}

func FindRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
        return dir, err
    }

	for dir != "" {

		paths = nil
		err := filepath.Walk(dir, visitFindRootDir)
	    if err != nil {
	        return dir, err
	    }

	    if len(paths) > 0 {
	    	break
	    }

	    dir = filepath.Dir(dir)
	}
	
	if dir == "" {
		return dir, errors.New("couldnt find .git folder...")
	}

	return dir, nil
}

func visitFind(path string, f os.FileInfo, err error) error {

    if f.IsDir() || f.Name() == ".git" {
        return nil
    }

    paths = append(paths, path)
    return nil
}

func visitFindRootDir(path string, f os.FileInfo, err error) error {

    if f.IsDir() && f.Name() == ".git" {
        paths = append(paths, path)
    }

    return nil
}  