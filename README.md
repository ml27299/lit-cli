# Lit CLI

The Lit CLI can be used to build modular applications using multiple git submodules

## Install

- via `curl`
    ```
    curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh | sudo bash

## Usage

After install, run 
	```
	lit init

in the working directory of your project, this will create two files "lit.link.json" and "lit.module.json" as well as initialize any submodules in the working directory

### lit.link.json
Ex.
```json
[
	{
		"dest" : "src/some/path",
		"sources" : ["some/source/path/*", "some/source/path/somefile.txt"]
	},

	{
		"dest" : "src/some/other/path",
		"sources" : ["some/source/path/somefile.txt"]
	},
]
```

### lit.module.json
Ex.
```json
[
	{
		"dest" : "some/path",
		"repo" : "git@github.com:USERNAME/PACKAGE.git"
	}
]
```


Run lit init to install any new submodules
