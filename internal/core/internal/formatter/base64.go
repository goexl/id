package formatter

import (
	"encoding/base64"
	"strconv"

	"github.com/goexl/id/internal/core/internal/core"
)

type Base64 struct {
	id core.Valuer
}

func NewBase64(id core.Valuer) *Base64 {
	return &Base64{
		id: id,
	}
}

func (b *Base64) Format() string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(b.id.Value(), 10)))
}
