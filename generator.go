package id

// Generator 生成器
type Generator interface {
	// Next 下一个标识
	Next() (Value, error)

	// Parse 从数字转换为标识
	Parse(from uint64) Value
}
