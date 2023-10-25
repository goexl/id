package id

import "github.com/goexl/log"

type snowflakeBuilder struct {
	node   int
	logger log.Logger
}

func newSnowflakeBuilder(node int, logger log.Logger) *snowflakeBuilder {
	return &snowflakeBuilder{
		node:   node,
		logger: logger,
	}
}

func (sb *snowflakeBuilder) Build() Generator {
	return newDefaultGenerator(newSnowflake(sb.node, sb.logger))
}
