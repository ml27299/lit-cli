package parser
import(
	//"fmt"
	"io/ioutil"
	"github.com/ml27299/lit-cli/helpers/paths"
	"bytes"
	"strings"
	"os"
	"github.com/spf13/viper"
	"runtime"
)

func parseFile(path string) error {
	file_string, err := ioutil.ReadFile(path)	
	if err != nil {
		return err 
	}

	viper.SetConfigType("toml")
	_ = viper.ReadConfig(bytes.NewBuffer([]byte(file_string)))

	return nil
}

func ConfigLinkItems(path string, parse bool, root_dir string) ([]LinkItem, error) {
	var response []LinkItem

	if parse {
		parseFile(path)
	}

	var keys []string
	for _, key := range viper.AllKeys() {
		if key[:4] == "link"{
			keys = AppendUnique(keys, strings.Split(key, ".")[0])
		}
	}

	// var sources []string
	// var removeSources []string

	for _, key := range keys {

		var sources []string
		var removeSources []string

		dest := viper.GetString(key+".dest")
		if runtime.GOOS == "windows" {
			dest = strings.Replace(dest, "/", "\\", -1)
		}

		normalized_dest, err := paths.NormalizeWithRoot(dest, root_dir)
		if err != nil {
			return nil, err
		}
		
		sources_interfaces := viper.Get(key+".sources")
		// sources = nil
		// removeSources = nil
		
		for _, sources_interface := range sources_interfaces.([]interface{}) {
			if sources_interface.(string)[:1] == "!" {
				str := sources_interface.(string)
				removeSources = append(removeSources, str[1:len(str)])
			}else {
				sources = append(sources, sources_interface.(string))
			}
		}

		if runtime.GOOS == "windows" {
			for i, _ := range sources {
				sources[i] = strings.Replace(sources[i], "/", "\\", -1)
			}
			for z, _ := range removeSources {
				removeSources[z] = strings.Replace(removeSources[z], "/", "\\", -1)
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

func ConfigModules(path string, parse bool) ([]GitModule, error) {
	var response []GitModule

	if parse {
		parseFile(path)
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

func ConfigViaPath(dir string, root_dir string) (ParseInfo, error) {
	
	var response ParseInfo
	// root_dir, err := paths.FindRootDir()
	
	// if err != nil {
	// 	return response, err
	// }
	var (
		linkItems []LinkItem
		modules []GitModule
	)

	if _, err := os.Stat(dir+"/.litconfig"); err == nil {
		parseFile(dir+"/.litconfig")
		linkItems, err = ConfigLinkItems(dir+"/.litconfig", false, root_dir)
		modules, err = ConfigModules(dir+"/.litconfig", false)
	}else {
		return response, err
	}

	return ParseInfo{
		LinkItems: linkItems,
		GitModules: modules,
		Config: ConfigInfo{},
	}, nil
}	

func Config(root_dir string) (ParseInfo, error) {

	var response ParseInfo
	// dir, err := paths.FindRootDir()
	
	// if err != nil {
	// 	return response, err
	// }

	var (
		linkItems []LinkItem
		modules []GitModule
	)

	if _, err := os.Stat(root_dir+"/.litconfig"); err == nil {
		parseFile(root_dir+"/.litconfig")
		linkItems, err = ConfigLinkItems(root_dir+"/.litconfig", false, root_dir)
		modules, err = ConfigModules(root_dir+"/.litconfig", false)
	}else {
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