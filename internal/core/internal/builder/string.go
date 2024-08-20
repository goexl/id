package builder

import (
	"github.com/goexl/id/internal/core/internal"
	"github.com/goexl/id/internal/core/internal/core"
	"github.com/goexl/id/internal/core/internal/formatter"
)

type String struct {
	id        core.Valuer
	formatter internal.Formatter
}

func NewString(id core.Valuer) *String {
	return &String{
		id:        id,
		formatter: formatter.NewBase64(id),
	}
}

func (s *String) Base64() (string *String) {
	s.formatter = formatter.NewBase64(s.id)
	string = s

	return
}

func (s *String) Base76() (string *String) {
	s.formatter = formatter.NewBase76(s.id)
	string = s

	return
}

func (s *String) Build() internal.Formatter {
	return s.formatter
}
