package id

var (
	_              = Base64
	_              = Base76
	_ stringOption = (*stringOptionFormatter)(nil)
)

type stringOptionFormatter struct {
	formatter stringFormatter
}

func StringFormatter(formatter stringFormatter) *stringOptionFormatter {
	return &stringOptionFormatter{
		formatter: formatter,
	}
}

func Base64() *stringOptionFormatter {
	return StringFormatter(newStringFormatterBase64())
}

func Base76() *stringOptionFormatter {
	return StringFormatter(newStringFormatterBase76())
}

func (f *stringOptionFormatter) applyString(options *stringsOptions) {
	options.formatter = f.formatter
}
