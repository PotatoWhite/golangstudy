# 5.2 출력

- tlfwp
```go
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 329234.8278

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a%d b:%d f:%f\n", a, b, f)
}
```
 - 특이점은 Println의 경우 각 Param 마다 space를 넣어준다.
 - Println은 %v로 동작한다. float의 경우 %g 로 동작함

- %f 6자리 이하는 실숫값 초과는 지수형태 출력
- %q 특수문자를 그대로 출력

# 자릿수

- %05d : 5자리로 출력(빈공간에 0이 삽입됨)
- %-05d, %-5d : 빈칸에 0이 추가 되지 않음 좌측정렬


# 
- 함수의 Param을 그냥 쓰면 항상 값으로 동작한다. 
- 주소가 필요한 경우 주소를 사용하는 경우도 있다.
  ```
  Scanln(&a, &b)
  ```

