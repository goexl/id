package param

type Snowflake struct {
	Node uint16
}

func NewSnowflake() *Snowflake {
	return &Snowflake{
		Node: 1,
	}
}
