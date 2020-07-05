package main

import (
	"testing"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/pkg/textparse"
	"github.com/stretchr/testify/assert"
)

func TestParseMetrics(t *testing.T) {
	t.Run("shouldParseMetricWithTypeAndNoLabels", func(t *testing.T) {
		input := `# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 13`

		expectedMetrics := Metrics{
			"go_threads": {
				Name:   "go_threads",
				Type:   "gauge",
				Labels: []string{},
			},
		}

		promParser := textparse.NewPromParser([]byte(input))

		metrics := parseMetrics(promParser)

		assert.Equal(t, expectedMetrics, metrics)
	})

	t.Run("shouldParseMetricsWithTypeAndLabels", func(t *testing.T) {
		input := `# HELP http_request_duration_microseconds The HTTP request latencies in microseconds.
# TYPE http_request_duration_microseconds summary
http_request_duration_microseconds{handler="prometheus",quantile="0.5"} 865.578
http_request_duration_microseconds{handler="prometheus",quantile="0.9"} 1266.81
http_request_duration_microseconds{handler="prometheus",quantile="0.99"} 1266.81
http_request_duration_microseconds_sum{handler="prometheus"} 2132.388
http_request_duration_microseconds_count{handler="prometheus"} 2
# HELP http_request_size_bytes The HTTP request sizes in bytes.
# TYPE http_request_size_bytes summary
http_request_size_bytes{handler="prometheus",quantile="0.5"} 63
http_request_size_bytes{handler="prometheus",quantile="0.9"} 63
http_request_size_bytes{handler="prometheus",quantile="0.99"} 63
http_request_size_bytes_sum{handler="prometheus"} 126
http_request_size_bytes_count{handler="prometheus"} 2
# HELP http_requests_total Total number of HTTP requests made.
# TYPE http_requests_total counter
http_requests_total{code="200",handler="prometheus",method="get"} 2`

		expectedMetrics := Metrics{
			"http_request_duration_microseconds": {
				Name:   "http_request_duration_microseconds",
				Type:   "summary",
				Labels: []string{"handler", "quantile"},
			},
			"http_request_duration_microseconds_sum": {
				Name:   "http_request_duration_microseconds_sum",
				Labels: []string{"handler"},
			},
			"http_request_duration_microseconds_count": {
				Name:   "http_request_duration_microseconds_count",
				Labels: []string{"handler"},
			},
			"http_request_size_bytes": {
				Name:   "http_request_size_bytes",
				Type:   "summary",
				Labels: []string{"handler", "quantile"},
			},
			"http_request_size_bytes_sum": {
				Name:   "http_request_size_bytes_sum",
				Labels: []string{"handler"},
			},
			"http_request_size_bytes_count": {
				Name:   "http_request_size_bytes_count",
				Labels: []string{"handler"},
			},
			"http_requests_total": {
				Name:   "http_requests_total",
				Type:   "counter",
				Labels: []string{"code", "handler", "method"},
			},
		}

		promParser := textparse.NewPromParser([]byte(input))

		metrics := parseMetrics(promParser)

		assert.Equal(t, expectedMetrics, metrics)
	})
}

func TestAddMetric(t *testing.T) {
	metrics := Metrics{
		"oneMetric": &Metric{
			Name:   "oneMetric",
			Labels: []string{"label"},
		},
	}

	expectedMetrics := Metrics{
		"oneMetric": &Metric{
			Name:   "oneMetric",
			Labels: []string{"label"},
		},
		"anotherMetric": &Metric{
			Name: "anotherMetric",
		},
	}

	metrics.AddMetric("oneMetric")
	metrics.AddMetric("anotherMetric")

	assert.Equal(t, expectedMetrics, metrics)
}

func TestSetMetricType(t *testing.T) {
	metrics := Metrics{
		"oneMetric": &Metric{
			Name:   "oneMetric",
			Labels: []string{"label"},
		},
	}

	expectedMetrics := Metrics{
		"oneMetric": &Metric{
			Name:   "oneMetric",
			Labels: []string{"label"},
			Type:   "gauge",
		},
		"anotherMetric": &Metric{
			Name: "anotherMetric",
			Type: "gauge",
		},
	}

	metrics.SetMetricType("oneMetric", "gauge")
	metrics.SetMetricType("anotherMetric", "gauge")

	assert.Equal(t, expectedMetrics, metrics)
}

func TestSetMetricLabels(t *testing.T) {
	metrics := Metrics{
		"oneMetric": &Metric{
			Name:   "oneMetric",
			Labels: []string{"label"},
		},
	}

	expectedMetrics := Metrics{
		"oneMetric": &Metric{
			Name:   "oneMetric",
			Labels: []string{"overridenLabel"},
		},
		"anotherMetric": &Metric{
			Name:   "anotherMetric",
			Labels: []string{"someLabel", "anotherLabel"},
		},
	}

	metrics.SetMetricLabels("oneMetric", []string{"overridenLabel"})
	metrics.SetMetricLabels("anotherMetric", []string{"someLabel", "anotherLabel"})

	assert.Equal(t, expectedMetrics, metrics)
}

func TestExtractLabelsAndMetricName(t *testing.T) {
	labels := labels.Labels{
		labels.Label{Name: "__name__", Value: "something"},
		labels.Label{Name: "oneLabel", Value: "oneValue"},
		labels.Label{Name: "anotherLabel", Value: "anotherValue"},
	}

	expectedMetricName := "something"
	expectedLabelNames := []string{"oneLabel", "anotherLabel"}

	labelNames, metricName := extractLabelsAndMetricName(labels)

	assert.Equal(t, expectedLabelNames, labelNames)
	assert.Equal(t, expectedMetricName, metricName)
}

func TestMetricsDifference(t *testing.T) {
	t.Run("NoDiff", func(t *testing.T) {
		metrics := Metrics{
			"oneMetric": &Metric{
				Name: "oneMetric",
			},
		}

		anotherMetrics := Metrics{
			"oneMetric": &Metric{
				Name: "oneMetric",
			},
		}

		metricNamesDiff := metrics.Diff(anotherMetrics)

		assert.Equal(t, []string{}, metricNamesDiff)
	})

	t.Run("OneMetricNameInDiff", func(t *testing.T) {
		metrics := Metrics{
			"oneMetric": &Metric{
				Name: "oneMetric",
			},
			"anotherMetric": &Metric{
				Name: "anotherMetric",
			},
		}

		anotherMetrics := Metrics{
			"oneMetric": &Metric{
				Name: "oneMetric",
			},
		}

		metricNamesDiff := metrics.Diff(anotherMetrics)

		assert.Equal(t, []string{"anotherMetric"}, metricNamesDiff)
	})
}
