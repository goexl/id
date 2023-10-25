package id

import (
	"github.com/goexl/log"
)

var (
	_        = Logger
	_ option = (*optionLogger)(nil)
)

type optionLogger struct {
	logger log.Logger
}

func Logger(logger log.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
