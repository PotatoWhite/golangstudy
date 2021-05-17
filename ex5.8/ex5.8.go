package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	fmt.Printf("%d params\n", n)

	if err != nil {
		fmt.Println(err)
		// 일단 버퍼를 비워주기 위해서 읽어야 함
		stdin.ReadString('\n')
	} else {
		fmt.Println(a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	fmt.Printf("%d params\n", n)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(a, b)
	}
}
