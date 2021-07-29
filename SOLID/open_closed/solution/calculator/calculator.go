package calculator

//Use the Strategy Design Pattern

type Calculator struct {
	Operaton Operation //interface
}

func (c Calculator) Calculate(a int, b int) int {
	return c.Operaton.Execute(a, b)
}
