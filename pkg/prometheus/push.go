package prometheus

import (
	"fmt"
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
		err := pusher.Push()
		if config.printLog {
			if err != nil {
				logger.Errorf(fmt.Sprintf("Could not push metric %s to pushgateway", metric.job), err)
			} else {
				logger.Infof("push metric %s to pushgateway success", metric.job)
			}
		}
	}
}
