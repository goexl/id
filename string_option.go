package id

type (
	stringOption interface {
		applyString(options *stringsOptions)
	}

	stringsOptions struct {
		formatter stringFormatter
	}
)

func defaultStringFormatter() *stringsOptions {
	return &stringsOptions{
		formatter: newStringFormatterDefault(),
	}
}
