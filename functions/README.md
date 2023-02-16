# 함수

함수는 잘 알다시피 자주 사용되는 코드 블록을 하나로 묶어서 재사용할 수 있다.     

함수를 정의하는 방식은 코틀린과 유사하다.

아래 코드를 보자.
```go
//func add(a, b int) int {
func add(a int, b int) int {
	return a + b
}

// 반환값이 없는 void의 경우
func void(a, b string) {

}
```
위 코드에서 보면 같은 타입이 파라미터로 온다면 주석된 부분처럼 처리할 수 있다.

이 방식은 앞서 변수를 선언할 때도 배웠는데 함수에서도 이 방법을 사용할 수 있다.      

반환값이 없다면 `void`함수처럼 리턴 타입을 정의하지 않으면 된다.      

자바나 코틀린을 할때 일반적으로 함수나 메서드명은 첫 글자를 소문자로 정의한다.     

근데 go에서는 이 대소문자가 의미가 있다.       

이유는 다른 패키지에서 호출하기 위해서는 이 함수명의 앞글자가 대문자여야 한다.     

여러분들은 이미 앞서 `연산자`챕터를 통해서 이런 코드들을 많이 봤을 것이다.     

또한 `fmt`패키지에서 콘솔에 로그를 찍는 `Print`같은 함수들을 봤을텐데 전부 대문자로 시작하는 것을 알 수 있다.

## 함수의 Argument, 즉 인자는 Parameter로 복사된다.

일반적으로 함수를 호출할 때 함수에 넘기는 값을 Argument, 즉 인자값이라고 표현한다.

함수의 입장에서는 Parameter, 즉 매개변수이다.     

일단 함수를 호출할 떄 우리가 인자값을 넘기면 함수가 실행될 떄 이 인자값이 파라미터에 복사가 된다.      

이것은 자바도 마찬가지이다.

우리가 흔히 언어를 배우면서 알게 되는 부분이 있다.

[Call by Value vs Call BY Reference](https://github.com/basquiat78/call-by-value-vs-call-by-reference-)     

예전에 써 둔 글이긴 하지만 한번 읽어 볼 만 하다.

좀 더 좋은 글은 구글링을 통해서 확인해 보시기 바란다.

go는 자바와 달리 포인터가 존재한다.      

이 부분은 차후 포인터에서 다시 한번 다뤄볼 것이다.

## 함수는 여러 개의 값을 반환한다.     

자바를 하다보면 이런 건 불가능한건가? 라는 생각을 해본 적이 있다.      

코틀린은 `data class`, `Pair`, `Triple`이나 다음과 같이 리스트를 통해서 사용할 수 있다.

```kotlin
fun multiReturn(): List<Any> {
    return listOf("a", 1000, listOf(1,2,3))
}

fun main() {
	val (a, b, c) = multiReturn()
	println("$a $b $c")
}
```
이런 식으로 사용할 수 있다.

앞서 우리는 `연산자`챕터를 통해서 `big.Float`를 사용하면서 잠깐 맛을 보긴 했다.       

그렇다면 이런 함수를 한번 만들어 보자.

```go
package main

import (
	"fmt"
)

func multiReturn() (string, int, float64, bool) {
	return "first", 2, 3.3, true
}

func main() {

	a, b, c, d := multiReturn()
	fmt.Println(a, b, c, d)

}
```
함수를 정의할 때 리턴 타입을 순서대로 정의하다.      

그리고 `return`문에서 정의한 리턴 타입에 맞춰서 구분자 `,`로 순서대로 정의하면 고스란히 사용할 수 있다.

코틀린에서 우리는 언더스코어 `_`를 통해서 컬렉션 함수나 람다내에서 사용하지 않는 변수의 경우 사용하는 방법이 있다.

물론 앞서 `연산자`챕터에서도 본 적이 있다.

여러개의 값을 반환하는 함수를 호출할 때 특정 값만 필요하다면 다음과 같이

```go
package main

import (
	"fmt"
)

func multiReturn() (string, int, float64, bool) {
	return "first", 2, 3.3, true
}

func main() {
	_, intValue, _, _ := multiReturn()
	fmt.Println(intValue)
}
```
선택적으로 사용할 수 있다.    

개인적으로 위 방법이 가장 편하다고 생각하지만 다른 방식도 존재한다.     

함수에서 리턴 타입을 정의할 때 아예 변수명까지 명시하는 경우이다.

```go
package main

import (
	"fmt"
)

func multiReturn2() (first string, second int, third float64, fourth bool) {
	first = "first"
	second = 2
	third = 3.3
	fourth = true
	return
}

func main() {
	f, s, t, _ := multiReturn2()
	fmt.Println(f, s, t)
}
```
위와 같은 방식도 가능하다. 

이때는 리턴 타입의 변수와 함수내의 변수명을 같게 하면 된다.

그리고 마지막으로 값을 지정하지 않고 `return`문으로 마무리하는 방식이다.     

때에 따라서는 이 두 방법중 하나가 더 유리할 경우가 생긴다. 

따라서 상황에 맞게 사용하면 좋다.

# At a Glance

이제부터는 본격적으로 로직을 만드는데 가장 중요한 뼈대중 하나인 `제어 흐름`을 다뤄볼 것이다.     

`if`, `for`, `switch`를 다루는 것으로부터 시작한다.     

자바와는 다른 go만의 특징이 있기 때문에 이 부분은 좀 더 자세하게 다뤄보고자 한다.

