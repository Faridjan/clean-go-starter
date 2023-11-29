package logger

import "sync"

var (
	once sync.Once

	svcNameSingleton string
	loggerSingleton  Interface
)

func InitLogger(serviceName string) {
	svcNameSingleton = serviceName

	once.Do(func() {
		loggerSingleton = &logger{}
	})
}

func GetLogger() Interface {
	return loggerSingleton
}

type logger struct {
}

func (l *logger) Log(kv ...interface{}) error {
	Debug(kv...)
	return nil
}

func (l *logger) Debug(kv ...interface{}) {
	Debug(kv...)
}

func (l *logger) Info(kv ...interface{}) {
	Info(kv...)
}

func (l *logger) Warn(kv ...interface{}) {
	Warn(kv...)
}

func (l *logger) Success(kv ...interface{}) {
	Success(kv...)
}

func (l *logger) Error(kv ...interface{}) {
	Error(kv...)
}

func (l *logger) Fatal(kv ...interface{}) {
	Fatal(kv...)
}
