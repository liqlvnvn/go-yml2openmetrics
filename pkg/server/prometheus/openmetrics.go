package main

import (
	"net/http"

	"log"

	"github.com/liqlvnvn/go-yaml2openmetrics/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	gauge1 = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "currencies",
			Name:      "usd",
		})
	gauge2 = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "currencies",
			Name:      "eur",
		})
)

func ServeHTTP(cfg config.Config) {
	prometheus.NewRegistry()
	prometheus.MustRegister(gauge1)
	prometheus.MustRegister(gauge2)

	http.Handle(cfg.MetricsPath, promhttp.Handler())
	// http.Handle("/metrics", func() http.Handler {
	// 	return promhttp.InstrumentMetricHandler(
	// 		prometheus.DefaultRegisterer, promhttp.HandlerFor(prometheus.Gatherers{}, promhttp.HandlerOpts{}),
	// 	)
	// }())

	gauge1.Set(70)
	gauge2.Set(80)

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
