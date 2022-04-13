package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/push"
)

func pusher() {
	for _, metric := range MetricMaps {

		pusher := push.New("http://192.168.3.9:9091", "test_woodpecker").Collector(metric.collector)

		for k, v := range metric.groupkv {
			pusher = pusher.Grouping(k, v)
		}

		if err := pusher.Push(); err != nil {
			fmt.Printf("Could not push completion time to Pushgateway: %e", err)
		}
	}
}
