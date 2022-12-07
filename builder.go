package id

import (
	"github.com/goexl/simaqian"
)

type builder struct {
	logger simaqian.Logger
}

func newBuilder(options *options) *builder {
	return &builder{
		logger: options.logger,
	}
}

func (b *builder) Snowflake(node int) *snowflakeBuilder {
	return newSnowflakeBuilder(node, b.logger)
}

func (b *builder) Autoincrement() *autoincrementBuilder {
	return newAutoincrementBuilder(1)
}
