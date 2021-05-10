package parser

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type YAML map[interface{}]interface{}

func ParseYMLFile(data string) (YAML, error) {
	// t := make(map[interface{}]interface{})
	t := new(YAML)

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}
	fmt.Println(t)

	return *t, nil
}
