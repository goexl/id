package builder

import (
	"github.com/goexl/id/internal/core"
	"github.com/goexl/id/internal/internal"
	"github.com/goexl/id/internal/param"
)

type Autoincrement struct {
	params *param.Autoincrement
	id     *Id
}

func NewAutoincrement(id *Id) *Autoincrement {
	return &Autoincrement{
		params: param.NewAutoincrement(),
		id:     id,
	}
}

func (a *Autoincrement) From(from uint64) (autoincrement *Autoincrement) {
	a.params.From = from
	autoincrement = a

	return
}

func (a *Autoincrement) GeneratorBuild() core.Generator {
	return internal.NewAutoincrement(a.params.From)
}

func (a *Autoincrement) Build() (id *Id) {
	a.id.generator = a
	id = a.id

	return
}
