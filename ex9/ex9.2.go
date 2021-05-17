package main

import "fmt"

func ex9_1() {
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨 켠다")
	} else if temp <= 3 {
		fmt.Println("히터 켠다")
	} else if temp <= 18 {
		fmt.Println("나가자")
	} else {
		fmt.Println("덥다")
	}
}

func ex9_2() {
	var age = 22

	if age >= 10 && age <= 15 {
		fmt.Println("You are young")
	} else if age > 30 || age < 20 {
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best Age")
	}
}

func ex9_3() {
	var age = 22

	if age >= 10 && age <= 15 {
		fmt.Println("You are young")
	} else if age > 30 || age < 20 {
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best Age")
	}
}

func main() {
	ex9_1()
	ex9_2()
}
