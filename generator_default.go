package id

var _ Generator = (*defaultGenerator)(nil)

type defaultGenerator struct {
	executor executor
}

func newDefaultGenerator(executor executor) *defaultGenerator {
	return &defaultGenerator{
		executor: executor,
	}
}

func (dg *defaultGenerator) Next() *Id {
	return dg.executor.next()
}
