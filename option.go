package id

import "github.com/goexl/simaqian"

type (
	option interface {
		apply(options *options)
	}

	options struct {
		logger simaqian.Logger
	}
)

func defaultOptions() *options {
	return &options{
		logger: simaqian.Default(),
	}
}
