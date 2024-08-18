package id

import (
	"github.com/goexl/log"
)

type builder struct {
	logger log.Logger
}

func newBuilder(options *Options) *builder {
	return &builder{
		logger: options.logger,
	}
}

func (b *builder) Snowflake(node uint) *snowflakeBuilder {
	return newSnowflakeBuilder(node, b.logger)
}

func (b *builder) Autoincrement() *autoincrementBuilder {
	return newAutoincrementBuilder(1)
}
