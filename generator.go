package id

// Generator 生成器
type Generator interface {
	// Next 下一个
	Next() (Value, error)
}
