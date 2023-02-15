package main

import (
	"fmt"
)

func main() {

	// 변수 선언 방법

	// FM 가장 일반적인 방식
	var variableName int = 10
	fmt.Println(variableName)

	// 타입 추론이 가능하기에 타입 생략
	var str = "test"
	var intValue = 10
	fmt.Println(str, intValue)

	// 선언대입문 :=을 활용하면 var와 타입 생략이 가능하다.
	oStr := "test"
	oIntValue := 10

	fmt.Println(oStr, oIntValue)

	// 개벌적으로 선언하기
	var a int
	var b int
	// 만일 같은 타입이라면 다음과 같이 선언이 가능하다.
	// e.g. var a, b, c, d int
	var c, d int
	fmt.Println(a, b, c, d)
	var f, g float64
	fmt.Println(f, g)
	var t, n bool
	fmt.Println(t, n)
	var str1, str2 string
	fmt.Println(str1, str2)

	// 상수 선언
	const RED string = "red"
	fmt.Println(RED)
	//상수는 변경할 수 없기 때문에 컴파일 시점에서 에러가 난다.
	//RED = "yellow"

	const F = 1.5

	var aa int = F * 100
	fmt.Println(aa)

	floatValue := 3.14

	// 여기서 floatValue는 정해진 타입 float64이기 때문에 int타입과 맞지 않아서 에러가 발생한다.
	//var bb int = floatValue * 100
	// 위 코드를 주석처리하고 타입추론을 하게 만들어 보자.
	var bb = floatValue * 100
	fmt.Println(bb)
	fmt.Printf("%T \n", bb)
	fmt.Println("=================")
	var under uint8 = 0
	var calculate = under - 1
	fmt.Println(calculate)

	var over uint8 = 255
	var calculateOver = over + 1
	fmt.Println(calculateOver)
}
