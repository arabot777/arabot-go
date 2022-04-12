package env

import "github.com/arabot777/arabot-go/pkg/conf/meta"

// 定义应用基本信息
type metaEnv struct {
	// 应用所属业务平台
	platform string
	// 应用所属服务
	service string
	// 运行环境: dev/test/uat/prod
	env meta.AppEnv
	// 版本号
	version string
	// 日志地址
	logPath string
}

func (m *metaEnv) Platform() string {
	return m.platform
}

func (m *metaEnv) Service() string {
	return m.service
}

func (m *metaEnv) Env() meta.AppEnv {
	return m.env
}

func (m *metaEnv) Version() string {
	return m.version
}

func (m *metaEnv) LogPath() string {
	return m.logPath
}
