package main

import (
	"fmt"
	"sort"
)

func main() {
	example01()
	example02()
	example03()
	example04()
	example05()
	example06()
	example07()
	example08()
	example09()
	example10()
}

func example01() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("array", array)
	fmt.Println("slice", slice)

	array[1] = 100
	fmt.Println("array", array)
	fmt.Println("slice", slice)

	slice = append(slice, 500)
	fmt.Println("array", array)
	fmt.Println("slice", slice)
}

func example02() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3:4]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
}

// slice 복사
func example03() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append([]int{}, slice1...)

	slice2[2] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}

// slice 복사
func example04() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	slice2[2] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}

// slice 중간 element 삭제
func example05() {
	slice := []int{1, 2, 3, 4, 5}
	idx := 2

	for i := idx + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}

	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}

// slice 중간 element 삭제 - append
func example06() {
	slice := []int{1, 2, 3, 4, 5}
	idx := 2
	slice = append(slice[:idx], slice[idx+1:]...)

	fmt.Println(slice)
}

// slice 중간 element 삭제 - copy
func example07() {
	slice := []int{1, 2, 3, 4, 5}
	idx := 2
	copy(slice[idx:], slice[idx+1:])
	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}

// slice 중간 element 삽입 - append
func example08() {
	slice := []int{1, 2, 3, 4, 5}
	idx := 2
	value := 100
	slice = append(slice[:idx], append([]int{value}, slice[idx:]...)...)

	fmt.Println(slice)
}

// slice 중간 element 삽입 - copy
func example09() {
	slice := []int{1, 2, 3, 4, 5}
	idx := 2
	value := 100
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = value

	fmt.Println(slice)
}

// slice sort
func example10() {
	slice := []int{1, 2, 3, 4, 5}

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	fmt.Println(slice)
	sort.Ints(slice)
	fmt.Println(slice)
}
