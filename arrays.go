package main

import "fmt"

type User struct {
	username string
	password string
}

func newUser(username string) *User {
	user := User{
		username: "n/a",
		password: "n/a",
	}
	return &user
}

func main() {

	//Inferred length array
	ints := [...]int{1, 2, 3, 4, 5}
	strs := [...]string{"foo", "bar", "may", "august"}

	//Defined length array
	strsD := [5]string{"f5", "0x", "lr", "rg", "fs"}

	//Struct array
	users := [...]User{{"u1", "123"}, {"u2", "345"}, {"u3", "678"}}

	fmt.Println(ints)
	fmt.Println(strs)
	fmt.Println(strsD)
	fmt.Println(users)
}
