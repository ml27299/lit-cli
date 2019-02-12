package parser
import(
	"fmt"
	"encoding/json"
	"io/ioutil"
	"../paths"
	"bytes"
	"strings"
	"errors"
	"os"
	"github.com/spf13/viper"
)

func ConfigLinkItems(path string) ([]LinkItem, error) {
	var response []LinkItem
	file_string, err := ioutil.ReadFile(path)	
	if err != nil {
		return response, err 
	}

	// via_json := func (json_string []byte) ([]LinkItem, error) {
	// 	if !json.Valid([]byte(json_string)) {
	// 		return response, errors.New("link json is not valid...")
	// 	}

	// 	json.Unmarshal([]byte(json_string), &response)
	// 	return response, nil
	// }

	via_toml := func (toml_string []byte) ([]LinkItem, error) {
		viper.SetConfigType("toml")

		err := viper.ReadConfig(bytes.NewBuffer([]byte(toml_string)))
		if err != nil {
			return response, nil
		}

		var keys []string
		for _, key := range viper.AllKeys() {
			if key[:4] == "link"{
				keys = AppendUnique(keys, strings.Split(key, ".")[0])
			}
		}

		var sources []string
		for _, key := range keys {

			dest := viper.GetString(key+".dest")
			sources_interfaces := viper.Get(key+".sources")
			sources = nil

			for _, sources_interface := range sources_interfaces.([]interface{}) {
				sources = append(sources, sources_interface.(string))
			}

			response = append(response, LinkItem{
				Dest: dest,
				Sources: sources,
			})
		}

		return response, nil
	}

	// response, err = via_json(file_string)
	// if err != nil {
		response, err = via_toml(file_string)
	//}

	if err != nil {
		return response, err
	}

	return response, nil
}

func ConfigModules(path string) ([]GitModule, error) {
	var response []GitModule
	file_string, err := ioutil.ReadFile(path)	
	if err != nil {
		return response, err 
	}

	// via_json := func (json_string []byte) ([]GitModule, error) {
	// 	if !json.Valid([]byte(json_string)) {
	// 		return response, errors.New("link json is not valid...")
	// 	}

	// 	json.Unmarshal([]byte(json_string), &response)
	// 	return response, nil
	// }

	via_toml := func (toml_string []byte) ([]GitModule, error) {
		viper.SetConfigType("toml")

		err := viper.ReadConfig(bytes.NewBuffer([]byte(toml_string)))
		if err != nil {
			fmt.Println(err.Error())
			return response, err
		}

		var keys []string
		for _, key := range viper.AllKeys() {
			if key[:9] == "submodule" {
				keys = AppendUnique(keys, strings.Split(key, ".")[0])
			}
		}

		for _, key := range keys {

			dest := viper.GetString(key+".path")
			repo := viper.GetString(key+".url")
			name := strings.Replace(key, "submodule", "", 1)

			response = append(response, GitModule{
				Repo: repo,
				Dest: dest,
				Name: name,
			})
		}

		return response, nil
	}

	// response, err = via_json(file_string)
	// if err != nil {
		response, err = via_toml(file_string)
	//}

	if err != nil {
		return response, err
	}

	return response, nil
}

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

	if _, err := os.Stat(dir+"/.litconfig"); err == nil {
		linkItems, err = ConfigLinkItems(dir+"/.litconfig")
	}

	if err != nil {
		return response, err
	}

	if _, err := os.Stat(dir+"/.litconfig"); err == nil {
		modules, err = ConfigModules(dir+"/.litconfig")
	}

	if err != nil {
		return response, err
	}

	return ParseInfo{
		LinkItems: linkItems,
		GitModules: modules,
		Config: ConfigInfo{},
	}, nil
}	


func Config2() (ParseInfo, error) {
	
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