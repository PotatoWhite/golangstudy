package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	example01()
	example02()
	example03()
	// example04() // infinite loop
	example05() // continue & break
	// example06()
	// example07()
	example08()
	example09()
	example10()
	example11()
}

func example01() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ",")
	}

	fmt.Println("")
}

// 초기문 생략
func example02() {
	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("for문이 종료 되었습니다.")
	i := 0
	for ; i < 10; i++ {
		fmt.Print(i, ",")
	}

	fmt.Print(i)
	fmt.Println("")
}

// 초기문 후처리 생략 (while)
func example03() {
	example09()
	i := 0
	for i < 10 {
		fmt.Print(i, ",")
		i++
	}

	fmt.Print(i)
	fmt.Println("")
}

// 무한 루프
func example04() {
	i := 0
	for {
		time.Sleep(time.Second)
		fmt.Print(i, ",")
		i++
	}
}

// continue * break
func example05() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		if i == 6 {
			break
		}

		fmt.Printf("6 * %d = %d\n", i, 6*i)
	}

	fmt.Println("")
}

// continue & break 02
func example06() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("입력하세요")
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("숫자로 입력하세요")
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력하신 숫자는 %d 입니다 . \n", num)
		if num%2 == 0 {
			break // 짝수 검사
		}
	}

	fmt.Println("for문이 종료 되었습니다.")
}

// * 찍기
func example07() {

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("for문이 종료 되었습니다.")
}

// * 찍기
func example08() {

	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("for문이 종료 되었습니다.")
}

// * 찍기 숙제
func example09() {

	for i := 5; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("for문이 종료 되었습니다.")
}

// flag break
func example10() {

	a := 1
	b := 2

	found := false

	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Println(a, " * ", b, "=", a*b)

}

// label break
func example11() {
	a := 1
	b := 2

Outerfor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break Outerfor
			}
		}
	}

	fmt.Println(a, " * ", b, "=", a*b)
}
