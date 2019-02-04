package parser
import(
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"../paths"
	"errors"
	"../resources"
	"os"
)

func Config() (ParseInfo, error) {
	
	var response ParseInfo
	dir, err := paths.FindRootDir()
	
	if err != nil {
		return response, err
	}

	if _, err := os.Stat(dir+"/lit.link.json"); os.IsNotExist(err) {

		data, err := resources.Asset("lit.link.json")
		file, err := os.Create(dir+"/lit.link.json")

		defer file.Close()
		
		if err != nil {
			return response, err
		}

		_, err = file.Write(data)
		if err != nil {
			return response, err
		}
	}

	if _, err := os.Stat(dir+"/lit.module.json"); os.IsNotExist(err) {

		data, err := resources.Asset("lit.module.json")
		file, err := os.Create(dir+"/lit.module.json")

		defer file.Close()

		if err != nil {
			return response, err
		}

		_, err = file.Write(data)
		if err != nil {
			return response, err
		}
	}
	
	//config_json_string, err := ioutil.ReadFile(dir+"/lit.config.json")
	links_json_string, err := ioutil.ReadFile(dir+"/lit.link.json")
	modules_json_string, err := ioutil.ReadFile(dir+"/lit.module.json")
	if err != nil {
		return response, err 
	}

	// if !json.Valid([]byte(config_json_string)) {
	// 	return response, errors.New("config json is not valid...")
	// }

	if !json.Valid([]byte(links_json_string)) {
		return response, errors.New("link json is not valid...")
	}

	if !json.Valid([]byte(modules_json_string)) {
		return response, errors.New("module json is not valid...")
	}
	
	var (
		linkItems []LinkItem
		modules []GitModule
		//config ConfigInfo
	)

	//json.Unmarshal(config_json_string, &config)
	json.Unmarshal([]byte(links_json_string), &linkItems)
	json.Unmarshal([]byte(modules_json_string), &modules)

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
	
	for _, ele := range slice {
        if ele == i[0] {
            return slice
        }
    }

	return append(slice, i[0])
}