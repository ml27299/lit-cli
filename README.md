# Lit CLI

The Lit CLI can be used to build modular applications utilizing an architecture built around hard links and git submodules. This cli extends the git cli so that the commands are indistinguishable

## Install

- via `curl`
    ```
    curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh | sudo bash

## Init
`lit init`

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
	}
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


Run "lit init" to install any new submodules

## Add
`lit add`

This command extends the `git add` command, but does it for all submodules and main application. All commands that work with `git add` work with `lit add`

Ex. `lit add .`




