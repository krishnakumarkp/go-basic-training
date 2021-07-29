package pets

type Dog struct {
	Pet //n golang we don’t have inherintance, we have a more powerful tool that is the composition.
}

func (d Dog) Greet() string {
	return "Hello i am a dog and my name is " + d.Name
}
