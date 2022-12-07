package id

import "github.com/goexl/simaqian"

var (
	_        = Logger
	_ option = (*optionLogger)(nil)
)

type optionLogger struct {
	logger simaqian.Logger
}

func Logger(logger simaqian.Logger) *optionLogger {
	return &optionLogger{
		logger: logger,
	}
}

func (l *optionLogger) apply(options *options) {
	options.logger = l.logger
}
