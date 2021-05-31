package main

import (
	"fmt"
	"unsafe"
)

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	example01()
	example02()
	example03()
	example04()
	example05()
	example06()
	example07()
	example08()

}

func example01() {
	var house House

	house.Address = "서울시 강남구"
	house.Size = 28
	house.Price = 10
	house.Category = "아파트"

	fmt.Println(house)
	fmt.Printf("%v\n", house)
	fmt.Printf("주소:%s 사이즈:%d평 가격:%v억원 종류:%s\n", house.Address, house.Size, house.Price, house.Category)
}

// 기본값 출력
func example02() {
	var house House = House{
		Address:  "여기가 어디야?",
		Size:     50,
		Category: "주택",
	}

	fmt.Println(house)
}

//
func example03() {
	type User struct {
		Name string
		ID   string
		Age  int
	}

	type VIPUser struct {
		UserInfo User
		VIPLevel int
		Price    int
	}

	user := User{"감자", "potato", 23}
	vip := VIPUser{
		User{"당근", "carrot", 48},
		3,
		250,
	}

	fmt.Printf("유저:%s ID:%s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저:%s ID:%s 나이:%d 레벨:%d 가격: %d\n", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)
}

// Embed
func example04() {
	type User struct {
		Name string
		ID   string
		Age  int
	}

	type VIPUser struct {
		User
		VIPLevel int
		Price    int
	}

	user := User{"감자", "potato", 23}
	vip := VIPUser{
		User{"당근", "carrot", 48},
		3,
		250,
	}

	fmt.Printf("유저:%s ID:%s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저:%s ID:%s 나이:%d 레벨:%d 가격: %d\n", vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price)
}

// Overriding
func example05() {
	type User struct {
		Name string
		ID   string
		Age  int
	}

	type VIPUser struct {
		User
		Name     string
		VIPLevel int
		Price    int
	}

	user := User{"감자", "potato", 23}
	vip := VIPUser{
		User{"당근", "carrot", 48},
		"사과",
		3,
		250,
	}

	fmt.Printf("유저:%s ID:%s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저:%s ID:%s 나이:%d 레벨:%d 가격: %d 본명:%s\n", vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price, vip.User.Name)
}

// pointer of struct
func example06() {
	type User struct {
		Name string
		ID   string
		Age  int
	}

	type VIPUser struct {
		User
		Name     string
		VIPLevel int
		Price    int
	}

	user := User{"감자", "potato", 23}
	vip := VIPUser{
		User{"당근", "carrot", 48},
		"사과",
		3,
		250,
	}

	vip2 := &vip

	fmt.Printf("유저:%s ID:%s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저:%s ID:%s 나이:%d 레벨:%d 가격: %d 본명:%s\n", vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.Price, vip.User.Name)
	fmt.Printf("%p\n", &vip)
	fmt.Printf("%p\n", &vip2)
	fmt.Printf("%p\n", vip2)
	vip.Name = "바꿔"
	fmt.Printf("%v\n", vip)
	fmt.Printf("%v\n", *vip2)
}

// size of struct
func example07() {
	type User struct {
		Age   int32   // 4byte
		Score float64 // 8byte
	}

	var a int = 45

	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a))

	// unsafe.Sizeof(user) 의 결과가 16byteㅇ다. 이유는 64bit 컴퓨터의 연산을 위해 레지스터 크기를 기초로 정렬한다.
	// 이렇게 정렬하는 것을 Padding이라 한다.
}

// padding
func example08() {
	type User struct {
		A int8 // 1byte
		B int  // 8byte
		C int8 // 1byte
		D int  // 8byte
		E int8 // 1byte
	}
	// 40 byte

	type User1 struct {
		A int8  // 1byte
		B int32 // 8byte
		C int8  // 1byte
		D int32 // 8byte
		E int8  // 1byte
	}
	// 40 byte
	type User2 struct {
		A int8 // 1byte
		C int8 // 1byte
		E int8 // 1byte
		B int  // 8byte
		D int  // 8byte
	}

	fmt.Printf("%d\n", unsafe.Sizeof(User{}))
	fmt.Printf("%d\n", unsafe.Sizeof(User1{}))
	fmt.Printf("%d\n", unsafe.Sizeof(User2{}))

}
