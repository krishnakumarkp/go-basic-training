package users

// type Userdata map[int]string

// func AddUser(d Userdata, id int, name string) {
// 	d[id] = name
// }

// func GetUser(d Userdata, id int) string {
// 	return d[id]
// }

// now for some reason when refactoring your package code you decided to change the userdata to userlist
// type Userlist map[int]string

// func AddUser(d Userlist, id int, name string) {
// 	d[id] = name
// }

// func GetUser(d Userlist, id int) string {
// 	return d[id]
// }

// this will break all the existing use of the package so a better way to do that will be

type Userdata map[int]string

type Userlist = Userdata

func AddUser(d Userlist, id int, name string) {
	d[id] = name
}

func GetUser(d Userlist, id int) string {
	return d[id]
}
