package main

import "fmt"

func main() {
	example01()
	example02()
	example03()
	example04()
	// example05()
}

func example01() {

	var a int = 10
	var p *int
	p = &a

	fmt.Printf("%v %v\n", a, &a)
	fmt.Printf("%v %v %v\n", p, &p, *p)

	*p = 20

	fmt.Printf("%v %v\n", a, &a)
	fmt.Printf("%v %v %v\n", p, &p, *p)

	b := p
	c := *p
	fmt.Printf("%v %v %v\n", *b, b, &b)
	fmt.Printf("%v %v \n", c, &c)

}

func example02() {
	fmt.Println("--------------------------")
	var a int = 10
	var p1 = &a
	var p2 = &a
	var p3 = &a

	fmt.Printf("포인터변수 실제 값은 %p %p %p\n", p1, p2, p3)
	fmt.Printf("포인터변수 주소는 %p %p %p\n", &p1, &p2, &p3)
	fmt.Printf("포인터 값은 %v %v %v\n", *p1, *p2, *p3)
}

func example03() {
	fmt.Println("--------------------------")
	var a int = 10
	var p1 *int = &a
	var p2 **int = &p1

	fmt.Printf("%v %v %v\n", a, *p1, **p2)
	fmt.Printf("%v %v %v\n", &a, &p1, &*p2)
	fmt.Printf("%v %v %v\n", &a, p1, *p2)
}

func example04() {
	type Data struct {
		data int
	}

	data1 := Data{}
	data1.data = 10

	data2 := &Data{}
	data2.data = 20

	fmt.Printf("stack %v\n", data1)
	fmt.Printf("heap %v\n", data2)
}
