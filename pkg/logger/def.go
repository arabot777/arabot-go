package logger

type (
	LogLevel int8
	Format   int8
)

// log level
const (
	LevelDebug LogLevel = iota - 1
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// output format
const (
	FormatText Format = iota + 1
	FormatJSON
)

// unify key name so that we can get a unified log search experience
const (
	KeyTime       = "timestamp"
	KeyLevel      = "level"
	KeyCaller     = "location"
	KeyMessage    = "msg"
	KeyStacktrace = "estack"
	KeyBizName    = "biz"
	KeyName       = "logger"
)

// default config
const (
	DefaultColor      = false
	DefaultStack      = true
	DefaultCallerSkip = 0
	DefaultLevel      = LevelDebug
	DefaultFormat     = FormatText
	DefaultTimeLayout = "2006/01/02 15:04:05"
)
