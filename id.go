package id

import (
	"strconv"
)

// Id 编号
type Id uint64

func (i *Id) Value() uint64 {
	return uint64(*i)
}

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
	buffers := make([]byte, 0, 22)
	buffers = append(buffers, '"')
	buffers = strconv.AppendUint(buffers, uint64(*i), 10)
	buffers = append(buffers, '"')

	return buffers, nil
}

func (i *Id) UnmarshalJSON(data []byte) (err error) {
	if original, pe := strconv.ParseUint(string(data[1:len(data)-1]), 10, 64); nil != pe {
		err = pe
	} else {
		*i = Id(original)
	}

	return
}
