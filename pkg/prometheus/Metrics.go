package prometheus

import (
	"github.com/arabot777/arabot-go/pkg/conf/meta"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"strconv"
	"sync"
)

var (
	metricMaps = make(map[string]*metrics)
	groupsBase = make(map[string]string, 4)
	once       sync.Once
	config     *pusherConfig
)

type metrics struct {
	job        string
	metricType MetricType
	collector  prometheus.Collector
	groupsMap  map[string]string
	pusher     *push.Pusher
}

func InitMetrics(m meta.Meta) {
	groupsBase["env"] = strconv.Itoa(int(m.Env()))
	groupsBase["platform"] = m.Platform()
	groupsBase["service"] = m.Service()
	groupsBase["version"] = m.Version()
	config = config.init("http://192.168.3.9:9091")

	go func() {
		loopPusher()
	}()
}

func Close() {

}

func Record(job string, value float64, metricType MetricType, groupsMap map[string]string) {
	switch metricType {
	case MetricType_GAUGE:
		recordGauge(job, value, metricType, groupsMap)
	case MetricType_COUNTER:
		recordCounter(job, value, metricType, groupsMap)
	}
}

func recordGauge(job string, value float64, metricType MetricType, groupsMap map[string]string) {
	if metric, ok := metricMaps[job]; ok {
		metric.groupsMap = groupsMap
		if gauge, ok := metric.collector.(prometheus.Gauge); ok {
			gauge.Set(value)
		}
	} else {
		once.Do(func() {
			metric := &metrics{
				job:        job,
				metricType: metricType,
				groupsMap:  groupsMap,
			}
			guage := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: job,
			})
			guage.Set(value)
			metric.collector = guage
			pusher := push.New(config.pushGatewayURL, job).Collector(guage)
			for k, v := range groupsBase {
				pusher = pusher.Grouping(k, v)
			}
			metric.pusher = pusher
			metricMaps[job] = metric
		})
	}
}

func recordCounter(job string, value float64, metricType MetricType, groupsMap map[string]string) {
	if metric, ok := metricMaps[job]; ok {
		metric.groupsMap = groupsMap
		if counter, ok := metric.collector.(prometheus.Counter); ok {
			counter.Add(value)
		}
	} else {
		once.Do(func() {
			metric := &metrics{
				job:        job,
				metricType: metricType,
				groupsMap:  groupsMap,
			}
			counter := prometheus.NewCounter(prometheus.CounterOpts{
				Name: job,
			})
			counter.Add(value)
			metric.collector = counter
			pusher := push.New(config.pushGatewayURL, job).Collector(counter)
			for k, v := range groupsBase {
				pusher = pusher.Grouping(k, v)
			}
			metric.pusher = pusher
			metricMaps[job] = metric
		})
	}
}
