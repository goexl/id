package internal

import (
	"time"
)

// Value 值
// 使用接口保证扩展性
type Value interface {
	String() string

	Time() time.Time

	Get() uint64
}
