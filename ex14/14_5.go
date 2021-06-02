package main

import "fmt"

type User struct {
	name string
	age  int
}

// escape analysing
// return 되는 &u를 분석하여 탈출하는 놈이 있으면 stack이 아니라 heap에 만든다.
func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
