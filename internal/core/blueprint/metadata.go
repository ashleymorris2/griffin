package blueprint

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Metadata struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Path        string `yaml:"-"`
}

func ReadMetadataFromFile(path string) (Metadata, error) {
	var meta Metadata

	data, err := os.ReadFile(path)
	if err != nil {
		return meta, err
	}

	if err := yaml.Unmarshal(data, &meta); err != nil {
		return meta, err
	}

	meta.Path = path
	return meta, nil
}

func ReadMetadataFromFiles(files []string) ([]Metadata, []error) {
	var metadata []Metadata
	var errs []error

	for _, file := range files {
		meta, err := ReadMetadataFromFile(file)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		metadata = append(metadata, meta)
	}

	return metadata, errs
}
