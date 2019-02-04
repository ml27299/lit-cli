package parser

import(
	"os"
	"strings"
	"path/filepath"
	"../paths"
)

type LinkItem struct {
	Dest string
	Sources []string
}

func (l *LinkItem) Expand() ([]Link, error) {
	var response []Link
	for _, og_source := range l.Sources {

		og_source, err := paths.Normalize(og_source)
		if err != nil {
			return response, err
		}

		if strings.Contains(og_source, "/*"){
			
			sources, err := paths.Find(filepath.Dir(og_source))
			if err != nil {
				return response, err
			}

			links, err := l.CreateLinks(sources, og_source)
			if err != nil {
				return response, err
			}

			response = append(response, links...)
			continue
		}

		link, err := l.CreateLink(og_source, og_source)
		if err != nil {
			return response, err
		}
		response = append(response, link)
	}

	return response, nil
}

func (l *LinkItem) CreateLinks(sources []string, og_source string) ([]Link, error) {
	var response []Link

	for _,source := range sources {

		link, err := l.CreateLink(source, og_source)
		if err != nil {
			return response, err
		}

		response = append(response, link)
	}

	return response, nil
}

func (l *LinkItem) CreateLink(source string, og_source string) (Link, error) {
	var (
		f os.FileInfo
		response Link
	)

	og_source_dir := filepath.Dir(og_source)
	source_dir := filepath.Dir(source)

	dest, err := paths.Normalize(l.Dest)
	f, err = os.Stat(source)
	if err != nil {
		return response, err
	}

	dest_ext := strings.Replace(source_dir, og_source_dir, "", -1)
	if dest_ext != "" {
		dest = dest+dest_ext
	}

	return Link{
		Dest: dest+"/"+f.Name(),
		Source: source,
	}, nil
}