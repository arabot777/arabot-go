package prometheus

type pusherConfig struct {
	pushGatewayURL string
}

func (push *pusherConfig) init(url string) *pusherConfig {
	return &pusherConfig{
		pushGatewayURL: url,
	}
}
