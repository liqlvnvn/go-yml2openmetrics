package promserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/liqlvnvn/go-yml2openmetrics/pkg/config"
	openmetrics "github.com/liqlvnvn/go-yml2openmetrics/pkg/openmetrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServeHTTP(cfg config.Config, openMetrixList []openmetrics.Gauge) {
	fmt.Printf("Starting the server... You can reach the output on %v%v%v\n", cfg.HostName, cfg.Port, cfg.MetricsPath)

	reg := prometheus.NewRegistry()

	for _, v := range openMetrixList {
		var gauge = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: v.Namespace,
				Name:      v.Name,
			})

		reg.MustRegister(gauge)
		// gauge.Set(v.Value.(float64))
		// if n, err := strconv.ParseFloat(v.Value, 64); err == nil {
		// 	gauge.Set(n)
		// }
		gauge.Set(float64(v.Value))
	}

	// Expose the registered metrics via HTTP.
	http.Handle(cfg.MetricsPath, promhttp.HandlerFor(
		prometheus.Gatherers{reg},
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
