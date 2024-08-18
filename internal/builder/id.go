package builder

import (
	"github.com/goexl/id/internal/param"
	"github.com/goexl/log"
)

type Id struct {
	params *param.Id
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
