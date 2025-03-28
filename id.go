package id

import (
	"time"
)

// Id 标识
// 使用接口保证扩展性
type Id interface {
	String() string

	Time() time.Time

	Value() uint64
}
