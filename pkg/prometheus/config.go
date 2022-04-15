package prometheus

type pusherConfig struct {
	pushGatewayURL string
	printLog       bool
}

func (push *pusherConfig) init(url string, printLog bool) *pusherConfig {
	return &pusherConfig{
		pushGatewayURL: url,
		printLog:       printLog,
	}
}
