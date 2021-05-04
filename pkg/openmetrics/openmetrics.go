package openmetrics

import (
	"fmt"

	"github.com/liqlvnvn/go-yml2openmetrics/pkg/parser"
)

// Temporary stub used in place of the library solution.
func GenerateOpenMetricsText(m parser.YAML) string {
	var openMetrixText string

	for _, v := range m.Currensies {
		openMetricsName := fmt.Sprintf("%s_%s", "currencies", v.Currency)
		openMetrixText += fmt.Sprintf("# HELP %s\n# TYPE %s gauge\n%s %s\n", openMetricsName, openMetricsName, openMetricsName, v.Value)
	}

	return openMetrixText
}
