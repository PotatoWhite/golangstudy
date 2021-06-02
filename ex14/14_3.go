package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func ChangeDataPointer(arg *Data) *Data {
	arg.value = 999
	arg.data[100] = 999

	return arg
}

func main() {

	var some Data

	ChangeData(some)
	fmt.Println(some.data[100])

	ChangeDataPointer(&some)
	fmt.Println(some.data[100])
}
