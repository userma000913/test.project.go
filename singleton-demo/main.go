package main

import "fmt"

type User struct {
	Name string
	Age  int
}

var u *User

func main() {

	New()
	fmt.Printf("%+v", u)

}

func New() *User {
	if u == nil {
		u = &User{
			Name: "a",
		}
	}
	return u
}
