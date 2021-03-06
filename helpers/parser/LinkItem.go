package parser

import(
	"strings"
	"path/filepath"
	"github.com/ml27299/lit-cli/helpers/paths"
	"runtime"
)

type LinkItem struct {
	Dest string
	NormalizedDest string
	Sources []string
	RemoveSources []string
}

func (l *LinkItem) Filter(sources []string, root_dir string) ([]string, error){
	var response []string

	for _, source := range sources {
		
		found := false
		for _, rm_source := range l.RemoveSources {

			rm_source, err := paths.NormalizeWithRoot(rm_source, root_dir)
			if err != nil {
				return response, err
			}

			if rm_source == source {
				found = true
				break
			}
		}

		if !found {
			response = append(response, source)
		}
	}

	return response, nil
}

func (l *LinkItem) FindMatchingSources(path string) ([]string, error) {
	
	var response []string
	for _, og_source := range l.Sources {
		if !strings.Contains(og_source, "*"){
			continue
		}

		response = AppendUnique(response, og_source)
	}

	return response, nil
}


func (l *LinkItem) Expand(root_dir string) ([]Link, error) {
	var response []Link
	for _, og_source := range l.Sources {

		og_source, err := paths.NormalizeWithRoot(og_source, root_dir)
		og_source_dir := filepath.Dir(og_source)
		if err != nil {
			return response, err
		}

		if strings.Contains(og_source, "*"){
			sources, err := paths.Find(filepath.Dir(og_source))
			if err != nil {
				return response, err
			}

			sources, err = l.Filter(sources, root_dir)
			if err != nil {
				return response, err
			}
			
			links, err := l.CreateLinks(sources, og_source_dir)
			if err != nil {
				return response, err
			}

			response = append(response, links...)
			continue
		}

		sources, err := l.Filter([]string{og_source}, root_dir)
		if err != nil {
			return response, err
		}

		link, err := l.CreateLink(sources[0], og_source_dir)
		if err != nil {
			return response, err
		}
		response = append(response, link)
	}

	return response, nil
}

func (l *LinkItem) CreateLinks(sources []string, og_source_dir string) ([]Link, error) {
	var response []Link

	for _,source := range sources {

		link, err := l.CreateLink(source, og_source_dir)
		if err != nil {
			return response, err
		}

		response = append(response, link)
	}

	return response, nil
}

func (l *LinkItem) CreateLink(source string, og_source_dir string) (Link, error) {

	source_dir := filepath.Dir(source)
	var source_slice []string

	if runtime.GOOS == "windows" {
		source_slice = strings.Split(source, "\\")
	}else {
		source_slice = strings.Split(source, "/")
	}
	source_name := source_slice[len(source_slice) - 1]

	dest := l.NormalizedDest
	dest_ext := strings.Replace(source_dir, og_source_dir, "", -1)
	if dest_ext != "" {
		dest = dest+dest_ext
	}

	if runtime.GOOS == "windows" {
		return Link{
			Dest: dest+"\\"+source_name,
			Source: source,
		}, nil
	}else {
		return Link{
			Dest: dest+"/"+source_name,
			Source: source,
		}, nil
	}
}