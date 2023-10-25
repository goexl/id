package id

import (
	"github.com/goexl/log"
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		logger log.Logger
	}
)

func defaultOptions() *options {
	return &options{
		logger: log.New().Apply(),
	}
}
