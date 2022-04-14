package main

import (
	"fmt"
	"github.com/arabot777/arabot-go/internal/example/helloworld/conf"
	"github.com/arabot777/arabot-go/pkg/logger"
	"github.com/arabot777/arabot-go/pkg/prometheus"
	"github.com/arabot777/arabot-go/pkg/signal"
	"time"
)

func main() {
	c, err := conf.ConfigureServiceEnv()
	if err != nil {
		panic(err)
	}
	// 尽可能在这里做个简单打印，并不计入日志
	// 方便运维同学查看传入的参数，调整错误的配置
	fmt.Println(c)

	err = logger.InitLogger(c.Meta)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	prometheus.InitMetrics(c.Meta)

	value := float64(5)
	for {
		prometheus.Record("test_hello_world", value, prometheus.MetricType_COUNTER, map[string]string{
			"customize": "good",
		})
		value += float64(1)
		time.Sleep(30 * time.Second)
	}

	signal.Wait()
}
