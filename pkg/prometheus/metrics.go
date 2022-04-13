package prometheus

import "github.com/prometheus/client_golang/prometheus"

var MetricMaps map[string]*Metric

type Metric struct {
	job        string
	metricType MetricType
	collector  prometheus.Collector
	groupsMap  map[string]string
}

func init() {
	//MetricMaps := make(map[string]Metric)
	//map[string]string{
	//	"a": "aa",
	//	"b": "bb",
	//}
}

func Record(job string, value float64, metricType MetricType, groupsMap map[string]string) {
	switch metricType {
	case MetricType_GAUGE:
		recordGauge(job, value, metricType)
	}
}

func recordGauge(job string, value float64, metricType MetricType) {
	if metric, ok := MetricMaps[job]; ok {
		if gauge, ok := metric.collector.(prometheus.Gauge); ok {
			gauge.Set(value)
		}
	} else {
		metric := &Metric{
			job:        job,
			metricType: metricType,
		}
		guage := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: job,
		})
		guage.Set(value)
		metric.collector = guage
		MetricMaps[job] = metric
	}
}
