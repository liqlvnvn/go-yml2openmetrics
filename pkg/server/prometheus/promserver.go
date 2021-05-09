package promserver

import (
	"fmt"
	"net/http"
	"strconv"

	"log"

	"github.com/liqlvnvn/go-yml2openmetrics/pkg/config"
	"github.com/liqlvnvn/go-yml2openmetrics/pkg/openmetrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// var (
// 	gauge1 = prometheus.NewGauge(
// 		prometheus.GaugeOpts{
// 			Namespace: "currencies",
// 			Name:      "usd",
// 		})
// 	gauge2 = prometheus.NewGauge(
// 		prometheus.GaugeOpts{
// 			Namespace: "currencies",
// 			Name:      "eur",
// 		})
// )

func ServeHTTP(cfg config.Config, openMetrixList []openmetrics.Gauge) {
	fmt.Printf("Starting the server... You can reach the output on %v%v%v\n", cfg.HostName, cfg.Port, cfg.MetricsPath)

	reg := prometheus.NewRegistry()
	for _, v := range openMetrixList {
		// fmt.Println(v)
		var gauge = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: v.Namespace,
				Name:      v.Name,
			})
		reg.MustRegister(gauge)
		// gauge.Set(v.Value.(float64))
		if n, err := strconv.ParseFloat(v.Value.(string), 64); err == nil {
			gauge.Set(n)
		}
	}
	// reg.MustRegister(gauge2)
	// reg.MustRegister(gauge1)

	// Expose the registered metrics via HTTP.
	http.Handle(cfg.MetricsPath, promhttp.HandlerFor(
		prometheus.Gatherers{reg},
		// prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	// gauge2.Set(80)
	// gauge1.Set(70)

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
