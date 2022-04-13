package prometheus

type MetricType int32

const (
	MetricType_COUNTER   MetricType = 0
	MetricType_GAUGE     MetricType = 1
	MetricType_SUMMARY   MetricType = 2
	MetricType_UNTYPED   MetricType = 3
	MetricType_HISTOGRAM MetricType = 4
)

