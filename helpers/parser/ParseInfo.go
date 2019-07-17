package parser
import(
	//"fmt"
	"strings"
	"github.com/ml27299/lit-cli/helpers/paths"
)

type ParseInfo struct {
	LinkItems []LinkItem
	GitModules []GitModule
	Config ConfigInfo
}

func (p *ParseInfo) GetLinks(root_dir string) ([]Link, error) {

	var response []Link

	root_dir, err := paths.FindRootDir()
	if err != nil {
		return response, err
	}

	for _, item := range  p.LinkItems {

		links, err := item.Expand(root_dir)
		if err != nil {
			return response, err
		}

		response= append(response, links...)
	}

	return response, nil
}

func (p *ParseInfo) FindMatchingLinkItems(path string, root_dir string) ([]LinkItem, error) {

	var response []LinkItem
	for _, item := range p.LinkItems {

		dest, err := paths.NormalizeWithRoot(item.Dest, root_dir)
		if err != nil {
			return response, err
		}	

		if strings.Contains(path, dest) {
			response = append(response, item)
		}
	}

	return response, nil
}

func (p *ParseInfo) FindMatchingLinkItemsBySubmodule(submodule_path, newfilepath string, root_dir string) ([]string, error) {

	var response []string
	for _, item := range p.LinkItems {

		dest, err := paths.NormalizeWithRoot(item.Dest, root_dir)
		if err != nil {
			return response, err
		}

		if !strings.Contains(newfilepath, dest){
			continue
		}

		for _, source := range item.Sources {

			if !strings.Contains(source, "/*") {
				continue
			}

			source, err := paths.Normalize(source)
			if err != nil {
				return response, err
			}
			
			if source == submodule_path+"/*" {
				response = AppendUnique(response, item.Dest)
				break
			}
		}
	}

	return response, nil
}