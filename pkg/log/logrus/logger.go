package logrus

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/enamespace/tpl/pkg/log"
)

type Logger struct {
	*logrus.Logger
}

var (
	std = New()
)

func C(ctx context.Context) *logrus.Entry {
	return std.C(ctx)
}
func New() *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) C(ctx context.Context) *logrus.Entry {
	ctxKeys, ok := ctx.Value(log.LogContextKey).([]string)
	if !ok {
		ctxKeys = []string{}
	}
	dataCopy := make(logrus.Fields, len(ctxKeys))
	for _, key := range ctxKeys {
		val := ctx.Value(key)
		dataCopy[key] = val
	}

	return l.WithFields(dataCopy)
}
