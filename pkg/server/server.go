package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/liqlvnvn/go-yml2openmetrics/pkg/config"
)

func ServeHTTP(cfg config.Config, content string) {
	fmt.Printf("Starting the server... You can reach the output on %v%v%v\n", cfg.HostName, cfg.Port, cfg.MetricsPath)
	http.HandleFunc(cfg.MetricsPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, content)
	})

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
