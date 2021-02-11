package logger

type Context interface {
	Log() *Logger
}
