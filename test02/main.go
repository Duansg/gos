package main

import "fmt"

type User struct {
	name string
}

func (u *User) PrintName(age int8) {
	fmt.Print(u.name)
	fmt.Print(age)
}

func main() {
	user := User{name: "Duansg"}
	user.PrintName(30)
}
