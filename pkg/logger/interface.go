package logger

type Interface interface {
	Log(kv ...interface{}) error

	Debug(kv ...interface{})
	Info(kv ...interface{})
	Warn(kv ...interface{})
	Success(kv ...interface{})
	Error(kv ...interface{})
	Fatal(kv ...interface{})
}
