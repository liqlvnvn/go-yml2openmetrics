package main

import (
	"fmt"
	"log"

	"github.com/liqlvnvn/go-yml2openmetrics/pkg/collector/file"
	"github.com/liqlvnvn/go-yml2openmetrics/pkg/config"
	"github.com/liqlvnvn/go-yml2openmetrics/pkg/openmetrics"
	"github.com/liqlvnvn/go-yml2openmetrics/pkg/parser"
	server "github.com/liqlvnvn/go-yml2openmetrics/pkg/server/prometheus"
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
	fmt.Println(om)
	// server.ServeHTTP(*cfg, om)
	server.ServeHTTP(*cfg)
}
