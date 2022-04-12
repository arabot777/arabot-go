package logger

import (
	"fmt"
	"os"
	"path/filepath"
)

// 创建开发环境日志
func NewDevLogger() Logger {
	return NewLogger(
		os.Stdout,
		WithColor(true),
		WithForamt(FormatText),
		WithLevel(LevelDebug),
	)
}

// 创建生产环境日志
func NewProdLogger(path, platform, service string) (Logger, error) {
	if _, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModeDir)
		if err != nil {
			fmt.Printf("fail to create dir for logger: %s", path)
			return nil, err
		}
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("fail to initialize logger because we can't open log file with given path: %s\n", path)
		return nil, err
	}
	return NewLogger(f,
		WithColor(false),
		WithForamt(FormatJSON),
		WithLevel(LevelInfo),
		WithField("platform", platform),
		WithField("service", service),
	), nil
}
