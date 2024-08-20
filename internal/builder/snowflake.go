package builder

import (
	"github.com/goexl/id/internal/core"
	"github.com/goexl/id/internal/internal"
	"github.com/goexl/id/internal/param"
)

type Snowflake struct {
	params *param.Snowflake
	id     *Id
}

func NewSnowflake(id *Id) *Snowflake {
	return &Snowflake{
		params: param.NewSnowflake(),
		id:     id,
	}
}

func (s *Snowflake) Node(node uint16) (snowflake *Snowflake) {
	s.params.Node = node
	snowflake = s

	return
}

func (s *Snowflake) GeneratorBuild() core.Generator {
	return internal.NewSnowflake(s.params.Node, s.id.params.Logger)
}

func (s *Snowflake) Build() (id *Id) {
	s.id.generator = s
	id = s.id

	return
}
