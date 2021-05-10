package openmetrics

import (
	"log"
	"reflect"

	parser "github.com/liqlvnvn/go-yml2openmetrics/pkg/parser/universal"
)

type Gauge struct {
	Namespace string
	Name      string
	Value     int
}

var gaugeList []Gauge

func GenerateOpenMetricsText(m parser.YAML) []Gauge {
	for namespace, v := range m {
		parseVal(v, namespace.(string))
	}

	return gaugeList
}

func parseVal(v interface{}, namespace string) {
	var gauge Gauge

	switch typ := reflect.TypeOf(v).Kind(); typ {
	case reflect.Slice:
		parseSlice(v.([]interface{}), namespace)
	case reflect.Map:
		if isMapWithValues(v.(parser.YAML)) {
			var value int

			var name string

			for nm, vl := range v.(parser.YAML) {
				if nm.(string) == "name" {
					name = vl.(string)
				} else if nm.(string) == "value" {
					value = vl.(int)
				}
			}

			gauge = Gauge{namespace, name, value}
			gaugeList = append(gaugeList, gauge)
		} else {
			parseMap(v.(parser.YAML), namespace)
		}
	default:
		log.Fatal("Error while parsing YAML file.")
	}
}

func isMapWithValues(m parser.YAML) bool {
	_, isMapContainsNameKey := m["name"]
	_, isMapContainsValueKey := m["value"]

	return isMapContainsNameKey && isMapContainsValueKey
}

func parseMap(m map[interface{}]interface{}, namespace string) {
	for k, v := range m {
		parseVal(v, namespace+(k.(string)))
	}
}

func parseSlice(slc []interface{}, namespace string) {
	for _, v := range slc {
		parseVal(v, namespace)
	}
}
