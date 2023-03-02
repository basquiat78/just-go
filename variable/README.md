# 변수

모든 언어를 배울 때 가장 먼저 알아야 하는 것은 변수일 것이다.

너무 당연하겠지만 각 언어마다 변수를 선언하는 방식은 비슷하면서도 다르다.

go는 코틀린이랑 묘하게 비슷하면서도 살짝 다르다.     

예전 코틀린에서도 [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)을 소개한 적이 있다.      

즉 사람이 글을 읽거나 코드를 읽는 방식에 대한 얘기이다.        

그렇다면 go는 어떻게 변수를 정의할까?

```go
var variableName int = 10
```
이런 형식이다. 자바보다는 확실히 코틀린에 좀 더 가까운 형태이다.

타입추론이 가능하기 때문에 다음과 같이 표현도 가능하다.

```go
package main

import (
	"fmt"
)

func main() {
	var str string = "test"
	var intValue = 10
	fmt.Println(str, intValue)

}
// variable> go run variable.go
// test 10
```

이것은 다음과 같이 좀 더 간략하게 표현이 가능하다.

```go
package main

import (
	"fmt"
)

func main() {
	str := "test"
	intValue := 10.00001
	fmt.Println(str, intValue)
}

// variable> go run variable.go
// test 10
```

선언대입문 `:=`은 선언과 대입을 한번에 한다는 의미를 지니고 있다.     

따라서 `var`키워드와 타입을 생략할 수 있다.      

## 숫자형 타입
자바와는 달리 go는 비트의 수로 표현한다.

예를 들면 기본적으로 go에서는 `int`라고 하면 컴퓨터의 비트에 따라 `int32` 또는 `int64`가 기본이다.

요즘도 32비트 컴퓨터를 쓰는지 모르겠지만 대부분 64비트일테니 `int64`이다.

자바처럼 `Long`, `Double`같은 타입으로 나누는 것이 아니라 이 비트를 통해서 구분한다.

또한 부호가 없는 정수라 해서 `uint`같은 타입도 존재한다.       

실수의 경우에는 `float32`, `float64`로 구분해서 사용한다.

특별히 go에는 별칭이 존재한다.    

그 중에 `byte`과 `rune`이 있는데 `byte`는 `uint8`, `rune`는 `int32`의 별칭이다.      

표를 통해서 한번 살펴보면 더 좋을 것 같다.

|     type     |       description       |                   range                    |
|:------------:|:-----------------------:|:------------------------------------------:|
|    uint8     |    1바이트 부호 없는 정수    |    0 ~ 255                                 |
|    uint16    |    2바이트 부호 없는 정수    |    0 ~ 65535                               |
|    uint32    |    4바이트 부호 없는 정수    |    0 ~ 4294967295                          |
|    uint64    |    8바이트 부호 없는 정수    |    0 ~ 18446744073709551616                |
|     int8     |    1바이트 부호 있는 정수    |    -128 ~ 127                              |
|     int16    |    2바이트 부호 있는 정수    |    -32768 ~ 32767                          |
|     int32    |    4바이트 부호 있는 정수    |    -2147483648 ~ 2147483647                |
|     int64    |    8바이트 부호 있는 정수    | -9223372036854775808 ~ 9223372036854775807 |
|    float32   |        4바이트 실수        |    IEEE-754 표준 32비트 실수                  |
|    float64   |        8바이트 실수        |    IEEE-754 표준64비트 실수                   |
|   complex64  |   8바이트 복소수(진수, 가수)  |    0 ~ 4294967295                         |
|  complex128  |  16바이트 복소수(진수, 가수)  |    0 ~ 18446744073709551616               |
|     byte     |        uint8의 별칭       |                                            |
|     rune     |        int32의 별칭       |                                            |
|     int      |      int32 또는 int64    |                                             |
|     uint     |     uint32 또는 uint64   |                                             |

## 타입의 기본 값은 무엇일까?

다음은 우리가 가장 많이 사용하게 될 타입별 기본값을 정의했다.      

참고로 go에서는 `null`대신 `nil`이라고 해서 정의되지 않은 메모리 주소를 나타내는 키워드가 존재한다.      

|              type             | default value |
|:-----------------------------:|:-------------:|
|    모든 정수타입 (int, uint계열)   |       0      |
| 모든 실수 타입(float, complex계열) |      0.0     |
|             bool              |     false     |
|            string             |  "" (빈 문자열) |
|             etc               |       nil     |

참고로 숫자형 타입의 경우에는 타입을 생략하게 되면 다음과 같은 방식을 따르게 된다.

만일 대입값이 정수라면 기본적으로 `int64`, 즉 `int`이고 실수면 `float64`로 결정된다.

## 변수 초기 선언
```go
package main

import (
	"fmt"
)

func main() {

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
}
```
위 코드를 실행하게 되면 앞서 배웠던 타입 기본값에 의거해서 초기값이 찍히게 될것이다.

## 상수 타입
상수는 잘 알듯이 변하지 않는 값을 의미한다.    

하지만 모든 타입을 상수로 선언할 수는 없다. 

상수로 사용될 수 있는 타입은 primitive type으로 다음과 같은 타입이 사용될 수 있다.

1. bool
2. 숫자형 타입 (uint, int, float, complex등)
3. string

이것은 자바스크립트와 상당히 유사한데 상수 선언은 `var`키워드 대신 `const`키워드를 사용한다.

```go
package main

import (
	"fmt"
)

func main() {

	// 상수 선언
	const RED string = "red"
	fmt.Println(RED)
	//상수는 변경할 수 없기 때문에 컴파일 시점에서 에러가 난다.
	//RED = "yellow"
}

```
상수를 사용해서 변하면 안되는 값을 설정해서 안전하게 사용할 수 있다.    

이유는 상수의 특징으로 비지니스 로직에서 상수를 실수로 바꿀 수 없기 때문이다.   

상수 선언시 타입을 생략해서 선언할 수 있다.     

보통 이런경우를 `타입없는 상수`라고 하는데 여기에는 특징이 하나 있다.      

바로 이 상수가 변수에 사용될 떄, 즉 복사가 되는 시점에 타입이 정해진다.

다음 코드를 보면 그 특이점이 발견된다.

```go
package main

import (
	"fmt"
)

func main() {

	const F = 1.5

	var aa int = F * 100
	fmt.Println(aa)

	floatValue := 1.5

	var bb int = floatValue * 100
	fmt.Println(bb)
}
```
여기서 특이한 점인 `F`를 상수로 선언했을 때 타입이 없기 때문에 1.5는 말 그대로 숫자라는 개념으로 동작하게 된다.   

그래서 `aa`라는 변수에는 계산된 결과 150이 대입이 된다. 

그런데 밑에 다음 코드를 보면 `floatValue`는 `float64`타입이 된다.      

여기서 `강 타입 언어`의 특징으로 들어온 변수 floatValue와 `int`타입이 서로 맞지 않아 오류가 발생한다.    

재미있는 것은 지금 `floatValue`의 타입으로 인해 다음과 같이 동작을 시키면 결과도 재미있다.

```go
package main

import (
	"fmt"
)

func main() {

	const F = 1.5

	var aa int = F * 100
	fmt.Println(aa)

	floatValue := 1.5

    // 여기서 floatValue는 정해진 타입 float64이기 때문에 int타입과 맞지 않아서 에러가 발생한다.
	//var bb int = floatValue * 100
	// 위 코드를 주석처리하고 타입추론을 하게 만들어 보자.
	var bb = floatValue * 100
    fmt.println(bb)
	fmt.Printf("%T", bb)
}
```
결과는 314가 나오지만 아래 코드를 통해 해당 변수의 타입을 찍어보면 `float64`로 찍힌다.     

이렇게 보면 숫자형 상수의 경우 `타입없는 상수`로 선언하게 되면 마치 코틀린의 `lazy`처럼 지연되는 효과처럼 보인다.

즉 `타입없는 상수`의 경우 그 값이 숫자형이라면 이런 특징을 사용해서 다양한 정수형 타입에 적절하게 사용할 수 있다.

go에는 `enum`을 따로 제공하지 않는다.

따라서 이 상수를 이용해서 enum처럼 사용할 수 있는 방법이 있는데 이것은 뒤에 따로 챕터를 마련할 것이다.

# Arithmetic Issues a.k.a Underflow & Overflow

지금은 `solidity`버전이 올라가면서 이런 부분이 자연스럽게 해소되었긴 하지만 그 이전에 `Underflow Attack` 이 버그를 이용한 공격이 있었다.      

그래서 SafeMath같은 라이브러리가 등장하기도 했는데 이런 버그는 go에도 존재한다.     

그렇다면 이 `Underflow`와 `Overflow`가 무엇일까 알아보자.

```go
package main

import (
	"fmt"
)

func main() {
	var under uint8 = 0
	var calculate = under - 1
	fmt.Println(calculate)
}
```

자 위 코드를 보고 `calculate`값이 어떻게 될지 한번 3초 정도 생각해 보자.   

일단 `uint8`의 범위는 위 표에서 보면 `0 ~ 255`이다.

그런데 이 최소값인 0에서 1를 빼면 어떤 일이 벌어질까?


원래대로라면 -1이 맞을 것이고 자바나 코틀린같은 언어라면 에러를 뱉어낼 것이다.

하지만 go에서는 놀랍게도 255가 찍힌다. 이것을 `Underflow`라고 한다.      

즉 이 버그를 이용해서 실행이 되면 안되는 코드가 실행 되버리는 것이다.

예를 들면 내가 가진 토큰이 100개이라고 하자. 

그런데 200개을 보낼려고 한다고 생각을 해보자.

만일 이 타입을 잘못 설정해서 위와 같이 `uint`를 타입으로 토큰의 발란스를 정했다면 해당 코드는 이런 코드를 가질 것이다.

참고로 solidity에서는 `uint`는 `uint256`과 같다.

```javascript
function withdraw(uint amount) {
    require(balances[msg.sender] - amount > 0);
    msg.sender.trasfer(amount)
    balances[msg.sender] -= amount
}
```
내가 가진 토큰이 100개인데 200개을 보낼려고 하면 위에서 `require`에서 걸러져야 한다.      

즉 자산이 부족하기 때문에 보낼 수 없어야 하는데 일단 200개를 보낸 것이다.

자신이 보유한 수량보다 더 많은 토큰을 보낸다는게 말이 될까?

놀랍게도 `100 - 200 > 0`에서 `Underflow`가 발생해버리기 때문에 `-100 > 0`라는 공식이 성립되지 않는다.

이것은 `uint범위의 최대값 > 0`이 되기 떄문에 true로 떨어지게 된다. 

게다가 보낸 사람의 토큰 발란스 역시 `uint범위의 최대값`이 되버리는 상황이 발생한다.

~~그리고 그것은 실제로 일어났습니다~~

이것을 이용해서 무한으로 토큰을 보낸 사건이 있었다.     

그렇다면 `Overflow`는 그 반대 개념이라고 생각하면 딱 맞다.

```go
package main

import (
	"fmt"
)

func main() {
	var over uint8 = 255
	var calculateOver = over + 1
	fmt.Println(calculateOver)
}

```
당신의 잔고는 0이 되는 놀라운 마법을 경험하게 될 것이다.       

따라서 uint, int계열의 정수형 타입을 사용할 경우, 특히 기본이 아닌 특정 비트의 계열을 사용하게 되면 이런 부분은 조심해야 한다.


# At a Glance

go의 변수 선언과 몇 가지 특징들을 살펴보았다.

다음으로 우리는 연산자에 대해서 알아 볼 것이다.