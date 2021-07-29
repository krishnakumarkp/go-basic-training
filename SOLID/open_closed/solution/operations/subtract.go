package operations

type Subtract struct {
}

func (m Subtract) Execute(a int, b int) int {
	return a - b
}
