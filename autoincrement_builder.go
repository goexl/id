package id

type autoincrementBuilder struct {
	from int64
}

func newAutoincrementBuilder(from int64) *autoincrementBuilder {
	return &autoincrementBuilder{
		from: from,
	}
}

func (ab *autoincrementBuilder) From(from int64) *autoincrementBuilder {
	ab.from = from

	return ab
}

func (ab *autoincrementBuilder) Build() Generator {
	return newDefaultGenerator(newAutoincrement(ab.from))
}
