package param

import (
	"github.com/goexl/log"
)

type Id struct {
	Logger log.Logger
}

func NewId() *Id {
	return &Id{
		Logger: log.New().Apply(),
	}
}
