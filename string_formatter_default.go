package id

import "strconv"

var _ stringFormatter = (*stringFormatterDefault)(nil)

type stringFormatterDefault struct{}

func newStringFormatterDefault() *stringFormatterDefault {
	return new(stringFormatterDefault)
}

func (sfd *stringFormatterDefault) format(id *Id) string {
	return strconv.FormatInt(int64(*id), 10)
}
