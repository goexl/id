package core

// Generator 生成器
type Generator interface {
	// Next 下一个
	Next() Id
}
