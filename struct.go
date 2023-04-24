package main

import "fmt"

type User struct {
	username string
	password string
}

func (user User) display() {
	fmt.Println("Username: ", user.username)
	fmt.Println("Password: ", user.password)
}

func newUser(username string) *User {
	user := User{
		username: "n/a",
		password: "n/a",
	}
	return &user
}

func (user User) getUsername() string {
	return user.username
}

func (user User) getPassword() string {
	return user.password
}

func main() {
	user := User{"deethesaint", "hehe"}

	fmt.Println(user.getUsername())
	fmt.Println(user.getPassword())

	user.display()

}
