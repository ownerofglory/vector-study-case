package logging

import "log"

type defaultLogger struct{}

func init() {
	InitDefaultLogger()
}

func InitDefaultLogger() {
	currentLogger = &defaultLogger{}
}

func (d *defaultLogger) Debug(args ...any) {
	log.Print(args...)
}

func (d *defaultLogger) Debugf(format string, args ...any) {
	log.Printf(format, args...)
}

func (d *defaultLogger) Info(args ...any) {
	log.Print(args...)
}

func (d *defaultLogger) Infof(format string, args ...any) {
	log.Printf(format, args...)
}

func (d *defaultLogger) Warn(args ...any) {
	log.Print(args...)
}

func (d *defaultLogger) Warnf(format string, args ...any) {
	log.Printf(format, args...)
}

func (d *defaultLogger) Error(args ...any) {
	log.Print(args...)
}

func (d *defaultLogger) Errorf(format string, args ...any) {
	log.Printf(format, args...)
}
