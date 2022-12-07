package id

import "github.com/goexl/simaqian"

type snowflakeBuilder struct {
	node   int
	logger simaqian.Logger
}

func newSnowflakeBuilder(node int, logger simaqian.Logger) *snowflakeBuilder {
	return &snowflakeBuilder{
		node:   node,
		logger: logger,
	}
}

func (sb *snowflakeBuilder) Build() Generator {
	return newDefaultGenerator(newSnowflake(sb.node, sb.logger))
}
