package log

import "context"

type contextKey int

const (
	LogContextKey contextKey = iota
)

type Logger interface {
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	C(ctx context.Context) Logger
}

const (
	NONE = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
)
