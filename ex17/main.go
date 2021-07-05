package main

import "fmt"

func main() {
	// example01()
	// example02()
	// example03()
	example04()
}

func example01() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("sliece is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice)
}

func example02() {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}
	var slice3 = make([]int, 3)

	if len(slice1) == 0 {
		fmt.Println("sliece is empty", slice1)
	}

	slice1[1] = 10
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func example03() {
	slice := []int{1, 2, 3}
	slice2 := append(slice, 4)
	slice3 := []int{1, 5: 2, 10: 3}

	fmt.Println(slice)
	fmt.Println(slice2)

	slice = append(slice2, 5)
	fmt.Println(slice)

	slice[1] = 100
	fmt.Println(slice)
	fmt.Println(slice2)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}

func example04() {
	array := [...]int{1, 2, 3, 4, 5}
	changeArray(array)
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}
	changeSlice(slice)
	fmt.Println(slice)

}

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}
