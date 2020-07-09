package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

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
	oldMetricsBytes, err := ioutil.ReadFile(oldMetricsFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	newMetricsBytes, err := ioutil.ReadFile(newMetricsFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	oldMetricsParser := textparse.NewPromParser(oldMetricsBytes)
	newMetricsParser := textparse.NewPromParser(newMetricsBytes)

	oldMetrics := parseMetrics(oldMetricsParser)
	newMetrics := parseMetrics(newMetricsParser)

	metricsDiff, _ := oldMetrics.Diff(newMetrics)

	fmt.Printf("Metrics Diff: \n\n")
	for _, metric := range metricsDiff {
		fmt.Println(metric)
	}
}

// Metric represents a prometheus metric
type Metric struct {
	Name   string
	Type   string
	Labels []string
}

// Metrics represents a collection of prometheus metrics
type Metrics map[string]*Metric

// SetMetricType sets the type of a metric given the metric name.
// It also creates the metric first if it does not exist
func (metrics Metrics) SetMetricType(metricName, metricType string) {
	metrics.AddMetric(metricName)
	metrics[metricName].Type = metricType
}

// SetMetricLabels sets the labels of a metric given the metric name.
// It also creates the metric first if it does not exist
func (metrics Metrics) SetMetricLabels(metricName string, metricLabels []string) {
	metrics.AddMetric(metricName)
	metrics[metricName].Labels = metricLabels
}

// AddMetric adds metric with the given metric name only
// if it does not exist. If it exists, it does not
// modify it
func (metrics Metrics) AddMetric(metricName string) {
	_, ok := metrics[metricName]
	if !ok {
		metrics[metricName] = &Metric{
			Name: metricName,
		}
	}
}

// MetricNameDiff represents the difference between
// two sets of metrics in terms of metric names
type MetricNameDiff []string

// MetricLabelDiff represents the difference between
// two metrics in terms of metric labels (label names),
// for the same metric name
type MetricLabelDiff struct {
	MetricName string
	LabelDiff  []string
}

// MetricLabelDiffs represents a list of MetricLabelDiff
type MetricLabelDiffs []MetricLabelDiff

// Diff finds the difference between the two metrics.
// What this means is that, what metric is present in
// metrics but not present in anotherMetrics
func (metrics Metrics) Diff(anotherMetrics Metrics) (MetricNameDiff, MetricLabelDiffs) {
	diff := []string{}

	for metricName := range metrics {
		_, metricExists := anotherMetrics[metricName]
		if !metricExists {
			diff = append(diff, metricName)
		}
	}

	return MetricNameDiff(diff), MetricLabelDiffs{}
}

func parseMetrics(parser textparse.Parser) Metrics {
	metrics := Metrics{}
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
			metrics.SetMetricType(string(metricName), string(metricType))
		}

		if entry == textparse.EntrySeries {
			var labels labels.Labels
			parser.Metric(&labels)
			labelNames, metricName := extractLabelsAndMetricName(labels)
			metrics.SetMetricLabels(metricName, labelNames)
		}
	}

	return metrics
}

func extractLabelsAndMetricName(labels labels.Labels) ([]string, string) {
	metricName := ""
	labelNames := make([]string, 0, len(labels)-1)
	for _, label := range labels {
		if label.Name == "__name__" {
			metricName = label.Value
			continue
		}

		labelNames = append(labelNames, label.Name)
	}
	return labelNames, metricName
}

func (metrics Metrics) String() string {
	var completeMetrics strings.Builder

	for _, metric := range metrics {
		completeMetrics.WriteString(fmt.Sprintf("%s, %s, %v\n", metric.Name, metric.Type, metric.Labels))
	}

	return completeMetrics.String()
}
