package parser

import (
	"log"

	"gopkg.in/yaml.v2"
)

// YAML struct define structure to match yaml file.
type YAML struct {
	Currensies []Field `yaml:"currencies"`
}

type Field struct {
	Currency string `yaml:"name"`
	Value    string `yaml:"value"`
}

func ParseYMLFile(data string) (YAML, error) {
	var t YAML

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	return t, nil
}
