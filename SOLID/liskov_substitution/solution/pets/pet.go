package pets

type Pet struct {
	Name string
}

func (p Pet) Greet() string {
	return "Hello i am a pet and my name is " + p.Name
}
