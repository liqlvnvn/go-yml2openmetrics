package main

import (
	"log"

	"github.com/liqlvnvn/go-yaml2openmetrics/pkg/collector/file"
	"github.com/liqlvnvn/go-yaml2openmetrics/pkg/config"
	"github.com/liqlvnvn/go-yaml2openmetrics/pkg/openmetrics"
	"github.com/liqlvnvn/go-yaml2openmetrics/pkg/parser"
	"github.com/liqlvnvn/go-yaml2openmetrics/pkg/server"
)

func main() {
	// Read config file.
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new YAML file entity.
	f := file.NewYAMLFile(cfg.InputSource)

	// Get content of the YAML file.
	data, err := f.Get()
	if err != nil {
		log.Fatal(err)
	}

	parsed, err := parser.ParseYMLFile(string(data))
	if err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	om := openmetrics.GenerateOpenMetricsText(parsed)

	server.ServeHTTP(*cfg, om)
}
