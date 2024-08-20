package param

type Autoincrement struct {
	From uint64
}

func NewAutoincrement() *Autoincrement {
	return &Autoincrement{
		From: 0,
	}
}
