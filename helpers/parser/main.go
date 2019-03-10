package parser
import(
	"fmt"
	"io/ioutil"
	"github.com/ml27299/helpers/paths"
	"bytes"
	"strings"
	"os"
	"github.com/spf13/viper"
)

func ConfigLinkItems(path string) ([]LinkItem, error) {
	var response []LinkItem
	file_string, err := ioutil.ReadFile(path)	
	if err != nil {
		return response, err 
	}

	viper.SetConfigType("toml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(file_string)))
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
	var removeSources []string
	for _, key := range keys {

		dest := viper.GetString(key+".dest")
		normalized_dest, err := paths.Normalize(dest)
		if err != nil {
			return response, err 
		}

		sources_interfaces := viper.Get(key+".sources")
		sources = nil
		removeSources = nil

		for _, sources_interface := range sources_interfaces.([]interface{}) {
			if sources_interface.(string)[:1] == "!" {
				str := sources_interface.(string)
				removeSources = append(removeSources, str[1:len(str)])
			}else {
				sources = append(sources, sources_interface.(string))
			}
		}

		response = append(response, LinkItem{
			Dest: dest,
			NormalizedDest: normalized_dest,
			Sources: sources,
			RemoveSources: removeSources,
		})
	}

	return response, nil
}

func ConfigModules(path string) ([]GitModule, error) {
	var response []GitModule
	file_string, err := ioutil.ReadFile(path)	
	if err != nil {
		return response, err 
	}

	viper.SetConfigType("toml")
	err = viper.ReadConfig(bytes.NewBuffer([]byte(file_string)))
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

func ConfigViaPath(dir string) (ParseInfo, error) {
	
	var response ParseInfo
	var (
		linkItems []LinkItem
		modules []GitModule
	)

	if _, err := os.Stat(dir+"/.litconfig"); err == nil {
		linkItems, err = ConfigLinkItems(dir+"/.litconfig")
	}else {
		return response, err
	}

	if _, err := os.Stat(dir+"/.litconfig"); err == nil {
		modules, err = ConfigModules(dir+"/.litconfig")
	}else {
		return response, err
	}

	return ParseInfo{
		LinkItems: linkItems,
		GitModules: modules,
		Config: ConfigInfo{},
	}, nil
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