# 변수

## 4.1 변수(Vari+able)
- 값이 다양할 수 있다.
- 컴퓨터적 용어 : 특정한 값을 갖고 있는다. 값을 갖고 있는 메모리 공간
  - 공간은 메모리에 있다. : 공간에 값을 저장하고 공간을 가르키는 것을 변수라 한다.

- 값(Data) : 컴퓨터화 되면 컴퓨터는 0,1 뿐이 모른다. 모든 것을 숫자로 판단한다.
- 프로그램 : 결국 데이터를 연산/조작 하는 일 (예, 화면에 그림을 그릴때 Video Memory의 각 픽셀의 RGB의 값을 변경하는 것이다.)
  - 이떄 변수를 통해 메모리에 접근해서 데이터를 읽어 오거나 쓰거나 한다.
- Camel 표기법 을 사용한다.
- '=' 대입연산자, 수학과 다르게 equal이 아니다. 
    ```golang
    a := 20
    b = a 
    ```
    - 해당 상황에서 a의 경우 값 b의 경우 메모리주소 로 Operator '=' 에 대해서 정의 된다.

## 4.2 변수 선언
- 사용을 위해서는 먼저 선언 해야한다.
  ```golang
  var a int = 10
  ```
  - var: 변수선언 키워드
  a : 변수명
  int : 변수 타입
  10 : 초깃값

## 4.3 변수에 대해 더 알아보기

### 4.3.1 변수의 4가지 속성
- 이름 : 이름을 통해서 프로그래머가 변수를 찾아간다. 컴파일시 메모리 주소로 대체된다.
  ```
  a := 3(값)
  MOV [0xffffffff], 3 
  ```

- 값 : 변수의 저장공간에 저장된 값
- 주소 : 연속된 메모리 중 해당 변수의 시작 주소
- 타입 : 실제 변수의 크기(길이, Size), sz?


## 숫자 타입

- uint8, uint16, uint32, unint64
- int8, int16, int32, nint64
- float32, float64
- byte : unit8의 alias
- rune : 문자타입 int32의 alias
- int : 숫자 Computer 마다 다르다.


## 그외 타입
- bool : true, false - default : false
- string : rune의 배열
- slice : 추후
- 구조체 : 필드들의 집합
- 포인터 : 추후
- nil : 

## 무조건 큰걸 안쓰는 이유는 메모리를 절약하기 위해


## 2진수 실수 표현
- float32 : 개략 소수부 7자리
- float64 : 개략 소수부 15자리
- 부호비트(1 bit) 지수부(8 bit) 소수부 (23 bit)