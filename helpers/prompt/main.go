package prompt
import(
	"github.com/chzyer/readline"
	"fmt"
	"strconv"
	. "github.com/ml27299/lit-cli/helpers"
	"strings"
	"github.com/ml27299/lit-cli/helpers/paths"
	//"gopkg.in/src-d/go-git.v4"
)

func PromptForUpdate() (string, error) {
	var response string

	Info("Would you like to update the cli to the latest version?")
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		response, err = rl.Readline()
		if err != nil {
			break
		}

		str := strings.TrimSpace(strings.ToLower(response))
		if err != nil {
			fmt.Println(err.Error())
		} else if str == "y" || str == "n" || str == "yes" || str == "no" {
			break
		}else {
			fmt.Println("Please enter Yy/Nn or yes/no")
		}
	}

	return response, nil
}

func PromptForInteractive(args []string, submodule *Module) (string, error) {
	var response string

	status, err := submodule.Status()

	if err != nil {
		return response, err
	}

	Info("Interactive mode for "+*&status.Path)
	Info("Below is the command that will be supplied to "+*&status.Path+", edit if not correct")

	rl, err := readline.NewEx(&readline.Config{
		Prompt: "> ",
		ForceUseInteractive: true,
	})
	rl.WriteStdin([]byte(strings.Join(args, " ")))
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		response, err = rl.Readline()
		if err != nil {
			break
		}

		break
	}
	
	return response, nil
}

func PromptForMultiSource(sources []string, newfilename string, newfilepath string) (int, error) {
	fmt.Println("There are multiple sources mapped to "+newfilepath)
	fmt.Println("Choose which source to add "+newfilename)

	var (
		response string
	)

	for index, source := range sources {
		fmt.Println("["+strconv.Itoa(index)+"] "+source)
	}

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		response, err = rl.Readline()
		if err != nil {
			break
		}

		index, err := strconv.Atoi(response)
		if err != nil {
			fmt.Println("Please enter a number for the corresponding sources")
		}else if index > len(sources) - 1 {
			fmt.Println("Please enter a number within range of the number of sources")
		} else {
			break
		}
	}

	responseInt, err := strconv.Atoi(response)
	if err != nil  {
		return responseInt, err
	}

	return responseInt, nil
}

func PromtForDest(dests []string, newfilename string, newfilepath string, submodule_path string) (int, error) {
	fmt.Println("There are multiple destinations pointing to "+newfilepath)
	fmt.Println("Choose which source to add "+newfilename)

	var (
		response string
	)

	for index, dest := range dests {

		dest, err := paths.Normalize(dest)
		if err != nil {
			return 0, err
		}

		newpath := strings.Replace(newfilepath, dest, "", -1)
		newpath = submodule_path+newpath

		fmt.Println("["+strconv.Itoa(index)+"] "+newpath)
	}

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		response, err = rl.Readline()
		if err != nil {
			break
		}

		index, err := strconv.Atoi(response)
		if err != nil {
			fmt.Println("Please enter a number for the corresponding destination")
		}else if index > len(dests) - 1 {
			fmt.Println("Please enter a number within range of the number of destinations")
		} else {
			break
		}
	}

	responseInt, err := strconv.Atoi(response)
	if err != nil  {
		return responseInt, err
	}

	return responseInt, nil
}

func PromptForNoDest(newfilename string, newfilepath string) (bool, error) {
	fmt.Println("Couldnt find a corresponding destination for "+newfilepath)
	fmt.Println("This means that you're trying to create a file in a linked directory, but there werent any matching sources to map to")
	fmt.Println("You can either add the file explicitly or implicitly (using /*) in lit.link.json to fix this")
	fmt.Println("If you wish to continue, "+newfilepath+"/"+newfilename+" will be created but will not link back to any submodule")
	fmt.Println("Do you wish to continue?")

	var (
		response string
	)

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		response, err = rl.Readline()
		if err != nil {
			break
		}

		response = strings.ToLower(response)
		if response != "yes" && response != "y"  && response != "no"  && response != "n" {
			fmt.Println("Please enter either Y/y/Yes/yes or N/n/No/no")
		}else {
			break
		}
	}

	responseBool := false
	if response == "yes" || response == "y"{
		responseBool = true
	}else if response == "no" || response == "n"  {
		responseBool = false
	}

	return responseBool, nil
}