package book

//A class should have only one reason to change.
//It means that your struct must only have one responsibility.
//It also means that when you create your struct, you should look at 3 things :

// 1 Is my method a behavior of my struct ? (Is this behavior the Single Responsability of my struct ?)
//	If yes, don’t change nothing it’s perfect, the only reason it would have to change, it’s if the behavior of your struct change.

type Book struct {
	Title   string
	Content string
}

func (b Book) GetTitle() string {
	return b.Title
}

func (b Book) GetContent() string {
	return b.Content
}
