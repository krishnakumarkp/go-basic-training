package calculator

type Operation interface {
	Execute(int, int) int
}
