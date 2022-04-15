package prometheus

import (
	"github.com/arabot777/arabot-go/pkg/conf/meta"
	"github.com/arabot777/arabot-go/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"strconv"
)

var (
	metricMaps = make(map[string]*metrics)
	groupsBase = make(map[string]string, 4)
	config     *pusherConfig
)

type metrics struct {
	job       string
	collector prometheus.Collector
	groupsMap map[string]string
	pusher    *push.Pusher
}

func InitMetrics(m meta.Meta) {
	groupsBase["env"] = strconv.Itoa(int(m.Env()))
	groupsBase["platform"] = m.Platform()
	groupsBase["service"] = m.Service()
	groupsBase["version"] = m.Version()
	config = config.init("http://192.168.3.9:9091", false)

	go func() {
		loopPusher()
	}()
}

func Close() {
	for _, metric := range metricMaps {
		pusher := metric.pusher
		_ = pusher.Delete()
		logger.Infof("record metric %s is closed", metric.job)
	}
}

func RecordGauge(job string, groupsMap map[string]string) prometheus.Gauge {
	var gauge prometheus.Gauge
	if metric, ok := metricMaps[job]; ok {
		metric.groupsMap = groupsMap
		gauge, _ = metric.collector.(prometheus.Gauge)
	} else {
		metric := &metrics{
			job:       job,
			groupsMap: groupsMap,
		}
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: job,
		})
		newPusher(job, metric, gauge)
	}
	return gauge
}

func RecordCounter(job string, groupsMap map[string]string) prometheus.Counter {
	var counter prometheus.Counter
	if metric, ok := metricMaps[job]; ok {
		metric.groupsMap = groupsMap
		counter, _ = metric.collector.(prometheus.Counter)
	} else {
		metric := &metrics{
			job:       job,
			groupsMap: groupsMap,
		}
		counter = prometheus.NewCounter(prometheus.CounterOpts{
			Name: job,
		})
		newPusher(job, metric, counter)
	}
	return counter
}

func newPusher(job string, metric *metrics, collector prometheus.Collector) {
	metric.collector = collector
	pusher := push.New(config.pushGatewayURL, job).Collector(collector)
	for k, v := range groupsBase {
		pusher = pusher.Grouping(k, v)
	}
	metric.pusher = pusher
	metricMaps[job] = metric
}
