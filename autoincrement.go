package id

type autoincrement struct {
	from uint64
}

func newAutoincrement(from uint64) *autoincrement {
	return &autoincrement{
		from: from,
	}
}

func (a *autoincrement) next() *Id {
	a.from++
	_id := Id(a.from)

	return &_id
}
