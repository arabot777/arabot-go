package env

import (
	"fmt"
	"github.com/arabot777/arabot-go/pkg/conf/meta"
	"github.com/arabot777/arabot-go/pkg/out"
	"os"
	"strconv"
)

const (
	NamePlatform = "PLATFORM"
	NameService  = "SERVICE"
	NameEnv      = "ENV"
	NameVersion  = "VERSION"
	NameLogPath  = "LOG_PATH"
)

type EnvReader struct{}

func NewReader() *EnvReader {
	return &EnvReader{}
}

func (r *EnvReader) GetMetaEnv() meta.Meta {
	m := &metaEnv{}
	var envStr string

	r.StringVar(&envStr, NameEnv, "")

	switch envStr {
	case "dev", "development", "debug":
		m.env = meta.EnvDev
	case "prod", "production", "release":
		m.env = meta.EnvProd
	case "test":
		m.env = meta.EnvTest
	case "uat", "pre":
		m.env = meta.EnvUat
	default:
		fmt.Printf("%s 初始化失败，请设置应用运行环境 %s 可选值: %s\n", out.Red.Add("[ERROR]"), out.Red.Add("ENV"), out.Yellow.Add("dev,development,debug/test/uat,pre/prod,production,release"))
		os.Exit(1)
	}

	r.StringVar(&m.platform, NamePlatform, "")
	if m.platform == "" {
		fmt.Printf("%s 初始化失败，请设置环境变量: %s\n", out.Red.Add("[ERROR]"), out.Red.Add("PLATFORM (业务线)"))
		os.Exit(1)
	}

	r.StringVar(&m.service, NameService, "")
	if m.service == "" {
		fmt.Printf("%s 初始化失败，请设置环境变量: %s\n", out.Red.Add("[ERROR]"), out.Red.Add("SERVICE (服务名)"))
		os.Exit(1)
	}

	r.StringVar(&m.version, NameVersion, "")
	if m.version == "" {
		fmt.Printf("%s 初始化失败，请设置环境变量: %s\n", out.Red.Add("[ERROR]"), out.Red.Add("VERSION (版本号)"))
		os.Exit(1)
	}

	r.StringVar(&m.logPath, NameLogPath, "")
	if m.logPath == "" {
		fmt.Printf("%s 初始化失败，请设置环境变量: %s\n", out.Red.Add("[ERROR]"), out.Red.Add("LOG_PATH (日志路径)"))
		os.Exit(1)
	}

	return m
}

// BoolVar 解析布尔类型的环境变量
// BoolVar parse bool system environment variable
func (e *EnvReader) BoolVar(p *bool, key string, val bool) {
	v, ok := os.LookupEnv(key)
	if !ok {
		*p = val
		return
	}
	if v != "true" && v != "false" && v != "1" && v != "0" {
		fmt.Printf("warning: os env [%s] should be a bool(true/false/1/0), but got: [%s], use default value\n", key, v)
		*p = val
		return
	}
	if v == "true" || v == "1" {
		*p = true
	} else {
		*p = false
	}
}

// IntVar 解析数字类型的环境变量
// IntVar parse int system environment variable
func (e *EnvReader) IntVar(p *int, key string, val int) {
	v, ok := os.LookupEnv(key)
	if !ok {
		*p = val
		return
	}
	v_, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("warning: os env [%s] should be an integer, but got: [%s], use default value\n", key, v)
		*p = val
		return
	}
	*p = v_
}

// StringVar 解析字符串类型的环境变量
// StringVar parse string system environment variable
func (e *EnvReader) StringVar(p *string, key string, val string) {
	v, ok := os.LookupEnv(key)
	if !ok {
		*p = val
		return
	}
	*p = v
}
