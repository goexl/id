package builder

import (
	"github.com/goexl/id/internal/core"
	"github.com/goexl/id/internal/param"
	"github.com/goexl/log"
)

type Id struct {
	params    *param.Id
	generator core.GeneratorBuilder
}

func NewId() *Id {
	return &Id{
		params: param.NewId(),
	}
}

func (i *Id) Logger(logger log.Logger) (id *Id) {
	i.params.Logger = logger
	id = i

	return
}

func (i *Id) Snowflake() *Snowflake {
	return NewSnowflake(i)
}

func (i *Id) Autoincrement() *Autoincrement {
	return NewAutoincrement(i)
}

func (i *Id) Build() core.Generator {
	return i.generator.GeneratorBuild()
}
