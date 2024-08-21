package formatter

import (
	"strconv"

	"github.com/goexl/gox"
	"github.com/goexl/id/internal/core/internal/core"
)

type Base76 struct {
	id core.Valuer
}

func NewBase76(id core.Valuer) *Base76 {
	return &Base76{
		id: id,
	}
}

func (b *Base76) Format() string {
	return gox.FormatUint(b.id.Value(), 10, strconv.FormatUint)
}
