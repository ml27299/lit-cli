# Lit CLI

The lit cli can be used to build modular applications utilizing an architecture built around hard links and git submodules. This cli extends the git cli so that the commands are indistinguishable. Hard links and git submodules are not a requirement for an application to use the lit cli, but the lit cli excels at managing applications that have submodules or hard linking

## Install

- via `curl`
    ```
    curl https://raw.githubusercontent.com/ml27299/lit-cli/master/install.sh | sudo bash


## Usage
`lit --help`

### Update lit to the latest version
`sudo lit update`

### Init
`lit init`

in the working directory of your project, this will create "lit.link.json", "lit.module.json" and ".gitignore", if the directory doesnt already have it, it will also initialize any submodules in the working directory (if the working directory has a .gitmodules file)

#### lit.link.json
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

#### lit.module.json
Ex.
```json
[
	{
		"dest" : "some/path",
		"repo" : "git@github.com:USERNAME/PACKAGE.git"
	}
]
```


Run `lit init` to install any new submodules

### Interactive mode
`lit --inter [command] [options]`__
`lit [command] [options] --inter`__
`lit [command] --inter [options]`__

The Interactive mode is useful when different options need to be passed depending on the submodule
Ex.

`lit commt -am "update something"` - this will send the message "update something" to all submodules and the main application

Sometimes you want to put different messages for different submodules, to accomplish this in lit run 

`lit --inter commit -am "update something"` - this will stop at every submodule and pop a prompt like

```console
Interactive mode for $SUBMODULE_PATH
Below is the command that will be supplied to $SUBMODULE_PATH, edit if not correct
> --message="upadate something" --all
```

Here you can edit the message
`> --message="update another something" --all`

### Specific submoldule
`lit --submodule $SUBMODULE_PATH [command] [options]`
or
`lit [command] [options] --submodule $SUBMODULE_PATH`
or 
`lit [command] --submodule $SUBMODULE_PATH [options]`

In lit you can throw any command to just one submodule using the `--submodule` flag
Ex.

lit --submodule $SUBMODULE_PATH commit -am "update"

### Link
`lit link`

Hard links files or file to a custom destination within the application (uses lit.link.json)

### Touch
`lit touch`

Since you may be hard linking a whole folder to various parts of your application, it can be a pain if you had to add 
a file to the source directory then relink. `lit touch` allows you to create a file anywhere in your application and will check if the directory of that file is being placed in has any source linking directories, if it does, it will create the file in the source directory and link it to the intended path

Ex. `lit touch ./src/some/path/newfile.txt`

If lit.link.json was configured to something like this
```json
[
	{
		"dest" : "src/some/path",
		"sources" : ["some/source/path/*"]
	},
]
```

then newfile.txt will be created in "some/source/path" and hard linked to "./src/some/path"


### Add
`lit add --help`

This command extends the `git add` command, but does it for all submodules and main application. All commands that work with `git add` work with `lit add`

Ex. `lit add .`

### Commit
`lit commit --help`

This command extends the `git commit` command, but does it for all submodules and main application. All commands that work with `git commit` work with `lit commit`

Ex. `lit commit -am "My first commit!"`

### Pull
`lit pull --help`

This command extends the `git pull` command, but does it for all submodules and main application. All commands that work with `git pull` work with `lit pull`

Ex. `lit pull origin master`

### Push
`lit push --help`

This command extends the `git push` command, but does it for all submodules and main application. All commands that work with `git pull` work with `lit push`

Ex. `lit push origin master`

### Merge
`lit merge --help`

This command extends the `git merge` command, but does it for all submodules and main application. All commands that work with `git pull` work with `lit merge`

Ex. `lit merge dev`

### Rebase
`lit rebase --help`

This command extends the `git rebase` command, but does it for all submodules and main application. All commands that work with `git rebase` work with `lit merge`

Ex. `lit merge dev`

### Checkout
`lit checkout --help`

This command extends the `git checkout` command, but does it for all submodules and main application. All commands that work with `git pull` work with `lit chckout`

Ex. `lit checkout dev`

### Extending
If anyone would like me to extend the lit cli to include more git commands, please submit an issue or pull request

