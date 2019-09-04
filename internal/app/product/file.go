package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/cecotw/strut-cli/internal/pkg/file"
	"github.com/ghodss/yaml"
)

// FileService Manages the product file
type FileService struct{}

// CreateFile Creates the product file in JSON or YAML
func (fs *FileService) CreateFile(model model) {
	var fileName = fmt.Sprintf("strut.%s", model.FileType.Extension)
	switch model.FileType {
	case file.Types.YAML:
		{
			yamlData, err := yaml.Marshal(model)
			if err != nil {
				log.Fatal(err)
			} else {
				err = ioutil.WriteFile(fileName, yamlData, 0644)
			}
		}
	case file.Types.JSON:
		{
			jsonData, err := json.Marshal(model)
			if err != nil {
				log.Fatal(err)

			} else {
				err = ioutil.WriteFile(fileName, jsonData, 0644)
			}
		}
	}
}

// ReadFile Loads the product file from the CWD
func (fs *FileService) ReadFile() (*Product, *error) {
	return nil, new(error)
}

// UpdateFile Updates the product file
func (fs *FileService) UpdateFile() {}

// AddApplication Adds an application to product file
func (fs *FileService) AddApplication() {}
