package prometheus

import (
	"github.com/arabot777/arabot-go/pkg/logger"
	"time"
)

func loopPusher() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			pusher()
		}
	}
}

func pusher() {
	for _, metric := range metricMaps {
		pusher := metric.pusher
		for k, v := range metric.groupsMap {
			pusher = pusher.Grouping(k, v)
		}
		if err := pusher.Push(); err != nil {
			logger.Errorf("Could not push metric to pushgateway", err)
		} else {
			logger.Infof("push metric to pushgateway success")
		}
	}
}
