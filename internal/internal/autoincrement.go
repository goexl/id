package internal

import (
	"github.com/goexl/id/internal/core"
)

type Autoincrement struct {
	from uint64
}

func NewAutoincrement(from uint64) *Autoincrement {
	return &Autoincrement{
		from: from,
	}
}

func (a *Autoincrement) Next() (id core.Id) {
	a.from++
	id = core.Id(a.from)

	return
}
