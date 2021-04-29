package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/liqlvnvn/go-yaml2openmetrics/pkg/config"
)

func ServeHTTP(cfg config.Config, content string) {
	http.HandleFunc(cfg.MetricsPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, content)
	})

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
