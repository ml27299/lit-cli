package paths

import (
	//"fmt"
	"os"
	"strings"
	"path/filepath"
	"errors"
)

var paths []string

func NormalizeWithRoot(str string, root_dir string) (string, error) {
	var response string
	
	current_path, err := os.Getwd()
	if err != nil {
        return response, err
    }

    var dir string
    if root_dir != current_path {

    	config_files, err := FindConfig(root_dir)
		if err != nil {
			return response, err
		}

    	for _, config_file := range config_files {
    		if filepath.Dir(config_file) == current_path {
    			dir = filepath.Dir(config_file)
    			break
    		}
    	}
    }else {
    	dir = root_dir
    }

    if dir == "" {
		return dir, errors.New("couldnt find dirctory for normalization...")
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

func Normalize(str string) (string, error) {
	var response string

	root_dir, err := FindRootDir()
	if err != nil {
		return response, err
	}
	
	current_path, err := os.Getwd()
	if err != nil {
        return response, err
    }

    var dir string
    if root_dir != current_path {

    	config_files, err := FindConfig(root_dir)
		if err != nil {
			return response, err
		}

    	for _, config_file := range config_files {
    		if filepath.Dir(config_file) == current_path {
    			dir = filepath.Dir(config_file)
    			break
    		}
    	}
    }else {
    	dir = root_dir
    }

    if dir == "" {
		return dir, errors.New("couldnt find dirctory for normalization...")
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

func FindConfig(path string) ([]string, error) {
	paths = nil

	err := filepath.Walk(path, visitFindConfig)
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

    if f.IsDir() || strings.Contains(f.Name(), ".git") {
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

func visitFindConfig(path string, f os.FileInfo, err error) error {

    if !f.IsDir() && f.Name() == ".litconfig" {
        paths = append(paths, path)
    }

    return nil
}  