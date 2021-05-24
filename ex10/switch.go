package main

import "fmt"

func example01() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a = 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	default:
		fmt.Println("a != 1,2,3", a)
	}
}

func example02() {
	temp := 9
	// AND 가 아닌 OR로 동작한다.
	switch true {
	case temp < 10, temp > 30:
		fmt.Println("나가지마")
		fallthrough
	case temp > 10, temp < 30:
		fmt.Println("나가놀아")
		fallthrough
	case temp > 100:
		fmt.Println("아 뜨거")
		fallthrough
	default:
		fmt.Println("모냐?")
	}
}

// example 03
func getMyAge() int {
	return 33
}

func example03() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	}

	// age는 초깃값을 설정한 범위에서만 접근이 가능하다.
	// fmt.Println(age)
}

// example 04
type ColorType uint8

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(colors ColorType) string {
	switch colors {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	example01()
	example02()
	example03()

	// example 04
	fmt.Println("My Favorite olor is", colorToString(getMyFavoriteColor()))
}

// Go 는 switch에 break를 적어도되고 안적어도 된다.
// break의 존재와 관계없이 case 실행 후 빠져나가는데.
// fallthrough를 통해 아래 case를 실행할 수 있다.
