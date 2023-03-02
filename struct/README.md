# 구조체

자바나 코틀린처럼 `go`에서는 클래스가 존재하지 않는다.     

하지만 `struct`, 즉 구조체로 이와 비슷한 개념이 존재한다.

아마도 이것은 c나 c++의 영향때문일 것이다.

어째든 클래스가 그랬듯이 이 구조체는 여러 다른 타입의 값들을 변수로 묶어주는 기능을 한다.

물론 메서드를 가질 수 있다.      

참고로 구조체는 자바의 클래스처럼 내부에 함수를 가질 수 없다. 

여기서 메서드와 함수에 대한 차이를 설명할 필요가 있는데 이 부분은 포인터를 다루는 챕터 이후에서 진행할 것이다.

또한 여기서 `go`가 지향하는 바가 무언인지도 살펴볼 것이다.

그렇다면 이 구조체를 어떻게 정의할 수 있는지 살펴보자.

가장 흔한 예제로 멤버로 예를 들어보자.     

멤버는 여러 변수를 자신의 상태값으로 가질 수 있다.     

그중에 간단하게 이름과 나이를 갖는 구조체를 하나 만들어보자.

```go
type Member struct {
    name string
    age int
}
```
아직 우리가 `go`에서 패키지에 대한 내용을 배운 적이 없다.    

`go`에서는 함수나 구조체를 정의할 때 해당 타입명이 대문자이냐 소문자이냐에 따른 차이점이 존재한다.

대문자인 경우에는 외부 패키지에서도 접근가능하게 된다. 일종의 `public`처럼 작동한다.     

하지만 소문자라면 외부로 공개되지 않는다. `private`이나 `protected`로 생각하면 좀 이해하기 쉽다.

게다가 자바처럼 멤버 변수의 초기값을 갖을 수 없다. 

이게 무슨 말일까?

```java
public class Member {

    //private String name;
    //private int age;
    private String name = "Anonymous";
    private int age = 25;
}
```
처럼 구조체 내의 멤버 변수의 값을 할당할 수 없다.

그리고 생성자가 없다.

자바를 예롤 들면 어떤 생성자도 만들지 않으면 기본적으로 인자값이 없는 기본 생성자를 내부적으로 자동생성한다.

따라서 만일 이런 방식을 사용하겠다고 한다면 이런 기능을 하는 함수를 정의해서 사용하는 방법을 택해야 한다.

먼저 구조체를 선언하는 방법부터 알아보자.     

기본 변수 선언과 크게 다르지 않다.

```go
package main

import (
	"fmt"
)

type Member struct {
	name string
	age  int
}

func main() {
	var member Member
	fmt.Println(member)

}
```
일단 그냥 구조체를 일반 변수 선언하듯 사용해서 사용해 보자.

앞서 배웠듯이 구조체의 선언된 멤버 변수들은 그에 해당하는 초깃값을 가지게 된다.

```
{ 0}
```
어떤 초기값도 선언하지 않는 구조체를 찍어보면 `string`의 초기값인 빈 값과 `int`의 초기값 0이 찍히는 것을 볼 수 있다.

## 모든 필드의 값을 초기화해보자.

다음과 같은 방식으로 각 멤버 변수의 값을 초기화할 수 있다.    

이때는 구조체내에 선안한 멤버 변수의 순서대로 처리한다.

```go
package main

import (
	"fmt"
)

type Member struct {
    name string
    age  int
}

func main() {
    member := Member{"basquiat", 25}
    fmt.Println(member)
}
```
여기서 만일 가독성을 위해서 또는 초기화할 멤버 변수가 많다면 다음과 같이 들여쓰기로 사용할 수 있다.

하지만 이때는 무조건 `trailing comma`, 즉 후행 컴마로 작성해야 한다.

```go
package main

import (
	"fmt"
)

type Member struct {
    name string
    age  int
}

func main() {
    member2 := Member{
        "basquiat",
        25, // 들여쓰기의 마지막 변수에는 무조건 후행 컴마 방식으로 ,를 찍어줘야 한다.
    }
    fmt.Println(member2)
}
```
다음과 같이 마지막 멤버 변수 초기값 이후에 무조건 `,`를 붙여줘야 한다.

물론 

```go
member2 := Member{
        "basquiat",
        25,}
```
이런 방식도 허용하지만 저장시에 플로그인에서 이 부분을 알아서 

```go
member2 := Member{
        "basquiat",
        25}
```
이렇게 변경시켜준다.


## Like Named Parameter
코틀린의 `네임드 파라미터`방식처럼 특정 멤버 변수만 초기값을 할당할 수도 있다.

다만 `=`이 아니라 마치 `json`처럼 `:`으로 매칭하게 된다.

```go
package main

import (
	"fmt"
)

type Member struct {
    name string
    age  int
}

func main() {
    member := Member{
        name: "basquiat",
    }
    fmt.Println(member)
}
// result
// {basquiat 0}
```
여기서도 위와 같은 방식이 통용된다.

즉 들여쓰기를 하는 경우에는 `후행 컴마`로 무조건 컴마를 붙여야 한다.

## struct in struct

다른 언어처럼 다른 객체, 여기서는 구조체를 포함할 수 있다.

```go
package main

import (
    "fmt"
)

type Member struct {
    name string
    age  int
}

type Warrior struct {
    memberInfo Member
    level      int
    ad         int
    ap         int
}

func main() {
    var warrior Warrior
    fmt.Println(warrior)
}
// result
// {{ 0} 0 0 0}
```

초기값을 할당한다면

```go
package main

import (
	"fmt"
)

type Member struct {
    name string
    age  int
}

type Warrior struct {
    memberInfo Member
    level      int
    ad         int
    ap         int
}

func main() {
    member := Member{"basquiat", 25}
    warrior := Warrior{
    memberInfo: member4,
        level:      50,
        ad:         25,
    }
    fmt.Println(warrior)
    fmt.Println(warrior1.memberInfo.name)
}

// result
// {{basquiat 25} 50 25 0}
// basquiat
```
이런 방식으로 선언할 수 있다.

구조체 내부의 구조체에 대한 접근 방식은 다른 언어와 똑같다.      

이런 방식은 다른 언어와 다를 바가 없다.      

근데 `go`에서는 `embedded field`라는 방식이 존재한다.

```go
package main

import (
	"fmt"
)

type Novice struct {
    name  string
    level int
}

type Rogue struct {
    Novice
    level int
    ad    int
    ap    int
}

func main() {
    novice := Novice{"basquiat", 10}
    rogue := Rogue{novice, 30, 50, 15}
    fmt.Println(rogue)
    fmt.Println(rogue.level)
    fmt.Println(rogue.name)
    fmt.Println(rogue.Novice.level)
}
```
`Rogue`구조체를 보면 `Novice`에 대한 변수명을 선언하지 않았다.

이 방식의 경우에는 '.'을 이용해 각 구조체의 멤버 변수에 접근할때 `Novice` 구조체의 멤버 변수를 접근할때도 `rogue.name`처럼 접근이 가능하다.

여기서 `level`이라는 멤버 변수가 `Novice`와 `Rogue`구조체 둘다 존재하는 경우가 있다.      

멤버 변수명이 중복된 경우에는 이 때는 명시적으로 `rogue.Novice.level`로 작성하면 된다.     

이 방식은 두 개의 구조체가 중복되는 멤버 변수가 없을 경우 아주 유용하다. 

즉 `rogue.Novice.name`이 아닌 `rogue.name`으로 접근이 가능해져 코드가 간결해 진다.

물론 중복되는 멤버 변수가 있다하더라도 `rogue.Novice.level`처럼 명시적으로 어느 구조체의 멤버 변수인지로 접근하면 되기 때문에 편의성을 가질 수 있는 방식이다.

## Memory Alignment And Padding

자바나 코틀린을 하면서 이런 생각을 해본적이 없었는데 구조체를 공부하다보면 `메모리 정렬`에 대한 개념을 접하게 된다.      

이것은 내부적으로 효율적으로 데이터를 읽기 위해서인데 `go`의 옵티마이져가 작동한다고 보면 된다.

예를 들면

```go
type Field struct {
    ad      int8
    ap      int8
    defense int64
    level   int8
    power   int64
}

type FieldSizeSort struct {
    defense int64
    power   int64
    ad      int8
    ap      int8
    level   int8
}
```
다음 두 개의 구조체를 살펴보자. 사실 이 두개의 구조체는 멤버 변수와 타입이 전부 똑같다.

다만 `Field`는 그냥 변수명의 알바벳 순서로 잡았고 `FieldSizeSort`는 타입의 크기를 내림차순으로 변수명의 알파벳 순서로 잡은 차이이다.

하지만 이 두 구조체가 차지하는 메모리 사이즈를 한번 측정해 보자.

```go
package main

import (
	"fmt"
	"unsafe"
)

type Field struct {
    ad      int8
    ap      int8
    defense int64
    level   int8
    power   int64
}

type FieldSizeSort struct {
    defense int64
    power   int64
    ad      int8
    ap      int8
    level   int8
}

func main() {
	var field Field
	var fieldSizeSort FieldSizeSort

	fmt.Println("field : ", unsafe.Sizeof(field))
	fmt.Println("fieldSizeSort : ", unsafe.Sizeof(fieldSizeSort))
}
// result
// field :  32
// fieldSizeSort :  24
```
`unsafe`패키지의 `Sizeof`함수를 통해서 값을 구하면 사이즈가 다른 것을 알 수 있다.     

이는 `메모리 정렬`과 `메모리 패딩`으로 인한 차이로 이와 좋은 글이 있어서 이 글로 대처한다.

~~영어라 미안합니다~~

[Golang struct size and memory optimisation](https://medium.com/techverito/golang-struct-size-and-memory-optimisation-b46b124f008d)

[Padding is hard](https://dave.cheney.net/2015/10/09/padding-is-hard)

사실 이 부분은 가독성을 염두 할 것인지 또는 구조체 내의 멤버 변수의 성격에 따라 정렬을 해서 작성할 것이냐에 대한 문제가 될 수 있다.     

그 정도의 메모리 낭비는 문제없다와 최적화를 하는게 중요하다의 대립각이 생길 수도 있을 것이다.      

## new 키워드를 이용해 구조체 선언

자바처럼 `new`키워드를 사용해 구조체를 할당하는 방법이 있다.

다만 이 경우에는 지금까지 배운 방식과는 좀 다르다.

먼저 이 키워드를 사용하면 멤버 변수의 초기값을 할당할 수 없다.

또한 이렇게 생성한 구조체를 콘솔에 찍어보면 해당 구조체는 스택이 아닌 힙 영역에 할당된 객체로 참조하고 있는 것을 볼 수 있다.

일단 거두절미하고 코드로 보자.

```go
package main

import (
	"fmt"
)

type Novice struct {
    name  string
    level int
}

func main() {
    newNovice := new(Novice)
    newNovice.name = "basquiat"
    fmt.Println(newNovice)
    fmt.Println(*newNovice)
}
// result
// &{basquiat 0}
// {basquiat 0}
```
일단 앞선 설명처럼 `new`키워드를 사용함과 동시에 멤버 변수의 초기값을 할당할 수는 없다.     

즉, `newNovice := new(Novice{name: "basquiat"})`이렇게 하고 싶지만 허용하지 않는다.

그리고 콘솔에 찍힌 내용을 보면 `&{basquiat 0}`처럼 앞에 `&`가 붙어 있는 것을 볼 수 있다.     

물론 `*`를 사용해 역참조를 하게 되면 지금까지 익숙하게 봤던 형식으로 찍히는 것을 볼 수도 있다.

즉 이 방식으로 구조체를 포인터로 받는 것을 알 수 있다.

# At a Glance

구조체에 대해서 알아 보았다.     

여기서 더 구체적인 내용을 담아야 하는데 그러기 위해서는 필연적으로 포인터에 대해서 먼저 알아야한다.

다음 챕터에서는 포인터를 다룰 예정이다.      

놓쳤던 설명을 이어 나갈 생각이다.
