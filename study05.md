# 상수는 값이 변하지 않는 것을 보장하기 위해 사용한다.
- 코드값으로 사용 : 코드란 순사에 의미를 부여하는 것이다.
  - ASCII 


# 상수는 Type이 없다.
- 특정 Type을 강제하는 경우 사용할 수도 있다.
- 대입연산자를 기준으로 상수는 좌변 값으로 사용할 수 없다.
  - 좌변은 공간의 의미로 사용되지만 상수는 값으로만 사용된다.
```go

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
    var a int = 3.14 * 100      // Compile시 314로 변경된다.
    var b int = FloatPI + 100 // 에러가 발생한다.

    fmt.Println(a)
    fmt.Println(b)
}
```

# 조건문

```
if 조건문 {

} else if 조건문 {

} else {

}
```