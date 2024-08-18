package id

type autoincrementBuilder struct {
	from uint64
}

func newAutoincrementBuilder(from uint64) *autoincrementBuilder {
	return &autoincrementBuilder{
		from: from,
	}
}

func (ab *autoincrementBuilder) From(from uint64) *autoincrementBuilder {
	ab.from = from

	return ab
}

func (ab *autoincrementBuilder) Build() Generator {
	return newDefaultGenerator(newAutoincrement(ab.from))
}
