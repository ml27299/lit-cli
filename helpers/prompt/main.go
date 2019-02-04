package prompt
import(
	"github.com/chzyer/readline"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"strconv"
	"strings"
	"../paths"
)

func PromptForSubmodule(submodules []*git.Submodule, newfilename string, newfilepath string) (int, error) {
	fmt.Println("There are multiple submodules in "+newfilepath)
	fmt.Println("Choose which submodule to add "+newfilename)

	var (
		response string
		paths []string
	)

	for _, submodule := range submodules {
		paths = append(paths, submodule.Config().Path)
	}

	for index, path := range paths {
		fmt.Println("["+strconv.Itoa(index)+"] "+path)
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
			fmt.Println("Please enter a number for the corresponding submodule")
		}else if index > len(submodules) - 1 {
			fmt.Println("Please enter a number within range of the number of submodules")
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