package id

func New(opts ...option) *builder {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return newBuilder(_options)
}
