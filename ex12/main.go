package main

import "fmt"

//  배열 : Random Access가 가장 빠른 자료구조

func main() {
	example01()
	example02()
	example02_01()
	example03()
	example04()
	example05_error()
	example06()
	example07()
	example08()
	example09_error()
	example10()

}

func example01() {
	fmt.Println("01 -------------")
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}

func example02() {
	fmt.Println("02 -------------")
	// 정적크기
	var n = [...]int{1, 2, 3}
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}

func example02_01() {
	fmt.Println("02_01 -------------")
	// 동적 배열 slice
	var n = []int{1, 2, 3}
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}

func example03() {
	fmt.Println("03 -------------")

	var f = [5]float64{1.0, 2.1, 3.2}
	for i := 0; i < len(f); i++ {
		fmt.Println(f[i])
	}
}

func example04() {
	fmt.Println("04 -------------")

	// n[1] = 10, n[3] = 30
	var n = [5]int{1: 10, 3: 30}
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}

func example05_error() {
	fmt.Println("05 -------------")
	// 배열의 크기는 변수를 사용할 수 없다.
	// x := 5
	// var n = [X]int{1: 10, 3: 30}
	// for i := 0; i < len(n); i++ {
	// 	fmt.Println(n[i])
	// }
}

func example06() {
	fmt.Println("06 -------------")
	nums := [...]int{10, 20, 30, 40, 50} // [5]int{10, 20, 30, 40, 50} 동일함
	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums)
	}
}

// range
func example07() {
	fmt.Println("07 -------------")
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	for index, value := range t {
		fmt.Println(index, value)
	}
}

// 배열 복사
func example08() {
	fmt.Println("08 -------------")

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{10, 20, 30, 40, 50}

	for _, v := range a {
		fmt.Println(v)
	}

	for _, v := range b {
		fmt.Println(v)
	}

	b = a

	for _, v := range b {
		fmt.Println(v)
	}
}

// 배열 복사
func example09_error() {
	fmt.Println("0 -------------")

	a := [5]int64{1, 2, 3, 4, 5}
	b := [5]int{10, 20, 30, 40, 50}

	for _, v := range a {
		fmt.Println(v)
	}

	for _, v := range b {
		fmt.Println(v)
	}

	// error type이 달라서 복제 안됨
	// b = a
}

// 다중배열 순회
func example10() {
	fmt.Println("10 -------------")

	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
