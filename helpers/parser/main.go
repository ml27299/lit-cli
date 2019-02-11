package parser
import(
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"../paths"
	"errors"
	"os"
)

func Config() (ParseInfo, error) {
	
	var response ParseInfo
	dir, err := paths.FindRootDir()
	
	if err != nil {
		return response, err
	}

	var (
		linkItems []LinkItem
		modules []GitModule
	)

	if _, err := os.Stat(dir+"/lit.link.json"); err == nil {

		links_json_string, err := ioutil.ReadFile(dir+"/lit.link.json")	
		if err != nil {
			return response, err 
		}

		if !json.Valid([]byte(links_json_string)) {
			return response, errors.New("link json is not valid...")
		}

		json.Unmarshal([]byte(links_json_string), &linkItems)
	}


	if _, err := os.Stat(dir+"/lit.module.json"); err == nil {

		modules_json_string, err := ioutil.ReadFile(dir+"/lit.module.json")	
		if err != nil {
			return response, err 
		}
	
		if !json.Valid([]byte(modules_json_string)) {
			return response, errors.New("module json is not valid...")
		}

		json.Unmarshal([]byte(modules_json_string), &modules)
	}

	return ParseInfo{
		LinkItems: linkItems,
		GitModules: modules,
		Config: ConfigInfo{},
	}, nil
}

func AppendUnique(slice []string, i ...string) []string {
	if len(i) > 1 {
	
		for _, val := range i {
			slice = AppendUnique(slice, val)
		}
		
		return slice
	}
		
	if len(i) == 0 {
		return slice
	}

	for _, ele := range slice {
        if ele == i[0] {
            return slice
        }
    }

	return append(slice, i[0])
}