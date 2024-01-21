package logging

//go:generate mockgen -package=logging -source=./logger.go -destination=./logger_mock.go
type Logger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
}

var currentLogger Logger

func InitLogger(logger Logger) {
	currentLogger = logger
}

func Debug(args ...any) {
	currentLogger.Debug(args...)
}

func Debugf(format string, args ...any) {
	currentLogger.Debugf(format, args...)
}

func Info(args ...any) {
	currentLogger.Info(args...)
}

func Infof(format string, args ...any) {
	currentLogger.Infof(format, args...)
}

func Warn(args ...any) {
	currentLogger.Warn(args...)
}

func Warnf(format string, args ...any) {
	currentLogger.Warnf(format, args...)
}

func Error(args ...any) {
	currentLogger.Error(args...)
}

func Errorf(format string, args ...any) {
	currentLogger.Errorf(format, args...)
}
