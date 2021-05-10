package file

import (
	"io/ioutil"
	"log"

	"github.com/liqlvnvn/go-yml2openmetrics/pkg/collector"
)

// Сreates a new entity.
func NewSource(filePath string) *collector.Source {
	return &collector.Source{Path: filePath}
}

type YAMLFile struct {
	filePath collector.Source
}

func NewYAMLFile(path string) *YAMLFile {
	return &YAMLFile{filePath: *NewSource(path)}
}

func (y *YAMLFile) Get() ([]byte, error) {
	data, err := ioutil.ReadFile(y.filePath.Path)
	if err != nil {
		log.Fatal("File reading error")
	}

	return data, err
}
