package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/textparse"
)

func main() {
	arguments := len(os.Args)
	if arguments != 3 {
		log.Fatal("usage: prom-exporter-metrics-diff <old-metrics-file-path> <new-metrics-file-path>")
	}
	oldMetricsFilePath := os.Args[1]
	newMetricsFilePath := os.Args[2]
	oldMetrics, err := ioutil.ReadFile(oldMetricsFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	newMetrics, err := ioutil.ReadFile(newMetricsFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	oldMetricsParser := textparse.NewPromParser(oldMetrics)
	newMetricsParser := textparse.NewPromParser(newMetrics)

	fmt.Printf("old metrics : \n")
	parseMetrics(oldMetricsParser)

	fmt.Printf("\n\n\nnew metrics : \n")
	parseMetrics(newMetricsParser)
}

func parseMetrics(parser textparse.Parser) {
	for true {
		entry, err := parser.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if entry == textparse.EntryType {
			metricName, metricType := parser.Type()
			fmt.Printf("Metric Name: %s, Metric Type: %s\n", metricName, metricType)
		}

		if entry == textparse.EntrySeries {
			var labels labels.Labels
			metric := parser.Metric(&labels)
			fmt.Printf("Metric : %s, Metric Labels: %v\n", metric, labels)
		}

		if entry == textparse.EntryUnit {
			metricName, metricUnit := parser.Unit()
			fmt.Printf("Metric Name: %s, Metric Unit: %s\n", metricName, metricUnit)
		}
	}
}
