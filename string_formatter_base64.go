package id

import (
	"encoding/base64"
)

var _ stringFormatter = (*stringFormatterBase64)(nil)

type stringFormatterBase64 struct{}

func newStringFormatterBase64() *stringFormatterBase64 {
	return new(stringFormatterBase64)
}

func (sfb *stringFormatterBase64) format(id *Id) string {
	return base64.StdEncoding.EncodeToString(id.Bytes())
}
