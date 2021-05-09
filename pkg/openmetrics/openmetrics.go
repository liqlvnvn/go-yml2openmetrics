package openmetrics

import (
	parser "github.com/liqlvnvn/go-yml2openmetrics/pkg/parser"
)

type Gauge struct {
	Namespace string
	Name      string
	Value     interface{}
}

// Temporary stub used in place of the library solution.
func GenerateOpenMetricsText(m parser.YAML) (gaugeList []Gauge) {
	for _, v := range m.Currensies {
		gaugeList = append(gaugeList, Gauge{"currensies", v.Currency, v.Value})
	}

	return
}
