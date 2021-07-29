package operations

type Add struct {
}

func (add Add) Execute(a int, b int) int {
	return a + b
}
