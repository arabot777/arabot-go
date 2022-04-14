package conf

import (
	"github.com/arabot777/arabot-go/pkg/conf/env"
	"github.com/arabot777/arabot-go/pkg/conf/meta"
)

// 配置文件
type Config struct {
	// 必须继承 Meta，包含应用元信息，防止10个服务10种配置，造成割裂
	meta.Meta
	Host    string
	Port    int
	RdsHost string
	RdsPort int
	RdsUser string
	RdsPass string
	RdsDB   string
	// mysql 可以汇聚成一组,这样初始化mysql的时候，只传入mysql的配置
	// 配置不多，直接写一起也不错
	// mysql 分组的配置:
	// Rds RdsConfig
}

// // mysql 配置汇聚成一组如下：
// type RdsConfig struct {
// 	RdsHost string
// 	RdsPort int
// 	RdsUser string
// 	RdsPass string
// 	RdsDB   string
// }

// 配置服务运行环境
func ConfigureServiceEnv() (*Config, error) {
	c := &Config{}

	reader := env.NewReader()

	// 检测和获取应用运行环境元信息: 包括业务线，服务名，版本号，release环境等
	c.Meta = reader.GetMetaEnv()

	// 这里开始就获取配置的参数
	reader.StringVar(&c.Host, "HOST", "")
	reader.IntVar(&c.Port, "PORT", 8080)

	return c, nil
}
