package main

import (
	"fmt"
	"math"
	"math/big"
)

// 인자 순서대로 좌변 우변을 비교한다.
// 자바와는 달리 함수로 전달하는 argumnet는 복사가 된다.
// 지금은 포인터를 배우지 않았지만 Cmp 함수는 포인터를 받기 때문에 타입에 포인터를 명시한다.
func gt(x, y *big.Float) bool {
	return x.Cmp(y) > 0
}

// 인자 순서대로 좌변 우변을 비교한다.
// 자바와는 달리 함수로 들어오는 파라미터는 복사가 된다.
// 지금은 포인터를 배우지 않았지만 Cmp 함수는 포인터를 받기 때문에 타입에 포인터를 명시한다.
func eq(x, y *big.Float) bool {
	return x.Cmp(y) == 0
}

func equalUseNextafter(x, y float64) bool {
	return math.Nextafter(x, y) == y
}

func main() {

	floatValue1 := 0.1
	floatValue2 := 0.2
	floatValue3 := 0.3

	// 0.3이라고 생각되지만 실제 값은 0.30000000000000004
	fmt.Println(floatValue1 + floatValue2)
	// 따라서 아래 비교연산자의 결과는 true가 아닌 false
	fmt.Println((floatValue1 + floatValue2) == floatValue3)

	// 다음과 같이 선언한다.
	// 코드를 따라가면 Float이라는 struct를 만들어서 객체화해서 사용한다.
	// 마치 wrapper class와 비슷하다.
	// 차후에 배우겠지만 go에서는 여러개의 값을 반환할 수 있다.
	bigF1, _ := new(big.Float).SetString("0.1")
	bigF2, _ := new(big.Float).SetString("0.2")
	bigF3, _ := new(big.Float).SetString("0.3")

	// struct 구조이기 때문에 Add라는 API를 이용해서 더한다.
	addValue := new(big.Float).Add(bigF1, bigF2)

	fmt.Println(addValue)

	// 마치 Cmp는 compare의 약자인듯 싶다.
	// 이것을 통해 0을 기준으로 좌변이 우변보다 크면 0보다 크고
	// 같으면 0을 반환한다.
	// 우변이 더 크다면 -1
	// 같다면 아래 표현식은 0을 반환할 때니 이것의 결과는 true가 떨어진다.
	fmt.Println(addValue.Cmp(bigF3) == 0)

	// 아직 함수와 포인터를 배우진 않았지만 이런 것은 함수로 따로 빼놔서 사용하는 것이 유리하다.
	// 인자의 순서대로 좌변, 우변으로 기준으로 한다.
	fmt.Println(gt(addValue, bigF3))
	fmt.Println(eq(addValue, bigF3))

	//Nextafter를 이용한 equalUseNextafter함수를 사용해 보자.
	fmt.Println(equalUseNextafter((floatValue1 + floatValue2), floatValue3))
}
