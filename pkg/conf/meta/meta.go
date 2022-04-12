package meta

type AppEnv int8

const (
	EnvDev AppEnv = iota + 1
	EnvTest
	EnvUat
	EnvProd
)

// 应用元信息.
type Meta interface {
	// 业务线: lbg/hrmt/lebesgue/bohrium/...
	Platform() string
	// 服务名: user/finance
	Service() string
	// 运行环境: dev/test/uat/prod
	Env() AppEnv
	// 版本号
	Version() string
	// 日志路径
	LogPath() string
}
