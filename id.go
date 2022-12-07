package id

import (
	"strconv"
)

// Id 编号
type Id int64

func (i *Id) String(opts ...stringOption) string {
	_options := defaultStringFormatter()
	for _, opt := range opts {
		opt.applyString(_options)
	}

	return _options.formatter.format(i)
}

func (i *Id) Bytes() []byte {
	return []byte(i.String())
}

func (i *Id) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(*i), 10)
	buff = append(buff, '"')

	return buff, nil
}

func (i *Id) UnmarshalJSON(data []byte) (err error) {
	if original, pe := strconv.ParseInt(string(data[1:len(data)-1]), 10, 64); nil != pe {
		err = pe
	} else {
		*i = Id(original)
	}

	return
}
