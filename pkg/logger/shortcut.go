package logger

import (
	"fmt"
	"github.com/arabot777/arabot-go/pkg/conf/meta"
	"os"
	"path/filepath"

	"go.uber.org/zap/zapcore"
)

const callerSkip = 2

var i *logger

func InitLogger(m meta.Meta) error {
	var err error
	switch m.Env() {
	case meta.EnvDev, meta.EnvTest:
		i = NewLogger(
			os.Stdout,
			WithColor(true),
			WithStack(false),
			WithForamt(FormatText),
			WithLevel(LevelDebug),
		)
	case meta.EnvProd, meta.EnvUat:
		if _, err := os.Stat(filepath.Dir(m.LogPath())); os.IsNotExist(err) {
			err = os.MkdirAll(m.LogPath(), os.ModeDir)
			if err != nil {
				fmt.Printf("fail to create dir for logger: %s", m.LogPath())
				return err
			}
		}
		f, err := os.OpenFile(m.LogPath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Printf("fail to initialize logger because we can't open log file with given path: %s\n", m.LogPath())
			return err
		}
		i = NewLogger(f,
			WithColor(false),
			WithForamt(FormatJSON),
			WithLevel(LevelInfo),
			WithField("platform", m.Platform()),
			WithField("service", m.Service()),
		)
	}
	return err
}

func Debugf(format string, v ...interface{}) {
	if ce := i.check(zapcore.DebugLevel, fmt.Sprintf(format, v...), callerSkip); ce != nil {
		ce.Write()
	}
}

func Infof(format string, v ...interface{}) {
	if ce := i.check(zapcore.InfoLevel, fmt.Sprintf(format, v...), callerSkip); ce != nil {
		ce.Write()
	}
}

func Warnf(format string, v ...interface{}) {
	if ce := i.check(zapcore.WarnLevel, fmt.Sprintf(format, v...), callerSkip); ce != nil {
		ce.Write()
	}
}

func Errorf(format string, v ...interface{}) {
	if ce := i.check(zapcore.ErrorLevel, fmt.Sprintf(format, v...), callerSkip); ce != nil {
		ce.Write()
	}
}

func Fatalf(format string, v ...interface{}) {
	if ce := i.check(zapcore.FatalLevel, fmt.Sprintf(format, v...), callerSkip); ce != nil {
		ce.Write()
	}
}

func DebugKV(msg string, kvs ...Entry) {
	fields := WrapKV(nil, false, kvs)
	if ce := i.check(zapcore.DebugLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func InfoKV(msg string, kvs ...Entry) {
	fields := WrapKV(nil, false, kvs)
	if ce := i.check(zapcore.InfoLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func WarnKV(msg string, kvs ...Entry) {
	fields := WrapKV(nil, false, kvs)
	if ce := i.check(zapcore.WarnLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func ErrorKV(msg string, kvs ...Entry) {
	fields := WrapKV(nil, false, kvs)
	if ce := i.check(zapcore.ErrorLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func FatalKV(msg string, kvs ...Entry) {
	fields := WrapKV(nil, false, kvs)
	if ce := i.check(zapcore.FatalLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func DebugErr(msg string, err error, kvs ...Entry) {
	fields := WrapKV(err, i.stack, kvs)
	if ce := i.check(zapcore.DebugLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func InfoErr(msg string, err error, kvs ...Entry) {
	fields := WrapKV(err, i.stack, kvs)
	if ce := i.check(zapcore.InfoLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func WarnErr(msg string, err error, kvs ...Entry) {
	fields := WrapKV(err, i.stack, kvs)
	if ce := i.check(zapcore.WarnLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func ErrorErr(msg string, err error, kvs ...Entry) {
	fields := WrapKV(err, i.stack, kvs)
	if ce := i.check(zapcore.ErrorLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func FatalErr(msg string, err error, kvs ...Entry) {
	fields := WrapKV(err, i.stack, kvs)
	if ce := i.check(zapcore.FatalLevel, msg, callerSkip); ce != nil {
		ce.Write(fields...)
	}
}

func Close() error { return i.Close() }

func IsInitialized() bool { return i != nil }

func GetLogger() Logger { return i }
