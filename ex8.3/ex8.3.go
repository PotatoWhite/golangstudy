package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

const (
	BitFlag1 uint = 1 << iota
	BitFlag2
	BitFlag3
	BitFlag4
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	}
}

func PrintColor(color int) {
	if color == Red {
		fmt.Println("빨")
	} else if color == Blue {
		fmt.Println("파")
	} else if color == Green {
		fmt.Println("녹")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)

	PrintColor(Blue)
}
