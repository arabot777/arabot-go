package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	ExamplePusherPush()
}

func ExamplePusherPush() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "woodpecker_difference_cloudserver_ali",
		Help: "Inconsistent data between clouderserver and ali.",
	})
	completionTime.Set(300)                                           // set可以设置任意值（float64）
	if err := push.New("http://192.168.3.9:9091", "test_woodpecker"). // push.New("pushgateway地址", "job名称")
										Collector(completionTime).                                  // Collector(completionTime) 给指标赋值
										Grouping("platform", "woodpecker").Grouping("env", "prod"). // 给指标添加标签，可以添加多个
										Push(); err != nil {
		fmt.Printf("Could not push completion time to Pushgateway: %e", err)
	}
}
