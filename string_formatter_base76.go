package id

import (
	"github.com/goexl/gox"
)

var _ stringFormatter = (*stringFormatterBase76)(nil)

type stringFormatterBase76 struct{}

func newStringFormatterBase76() *stringFormatterBase76 {
	return new(stringFormatterBase76)
}

func (sfb *stringFormatterBase76) format(id *Id) string {
	return gox.FormatIntd(int64(*id))
}
