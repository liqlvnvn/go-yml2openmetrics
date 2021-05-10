package parser

import (
	"log"

	"gopkg.in/yaml.v2"
)

type YAML map[interface{}]interface{}

func ParseYMLFile(data string) (YAML, error) {
	t := new(YAML)

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	return *t, nil
}
