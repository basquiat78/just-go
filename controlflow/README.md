# 제어 흐름
go는 단순한 면이 있다.     

예를 들면 `while`같은 제어문이 없다.

이것은 `for`문을 통해서 작성하게 되는데 어찌되었든 가장 중요한 뼈대이기 때문에 하나씩 일아가 보자.

# if
가장 기본적인 분기문은 바로 `if`이다.

어떤 조건을 만족하면 특정 로직을 태우는 방식이다.

이것은 대화로 풀면 바로 이해가 간다.

```
내가 달성치를 30프로이상 달성했다면, 랩을 할께.
그런데 30프로를 달성하지 못하면, 춤을 출께.
```

이런 일련의 대화 형식을 코드로 풀어보는것이 가장 좋을 것이다. 

위 내용을 토대로 `go`에서는 어떻게 사용하는지 먼저 살펴보자.

```go
package main

import (
	"fmt"
)

func rap() {
	fmt.Println("울려댔어~ 싸이렌~ 텅빈 길거리엔~ 도망치다 흘린 칼자루와 피가 흥건해~")
}

func dance() {
	fmt.Println("이것은 남자의 제로투~~~~")
}

func main() {
	//내가 달성치를 30프로이상 달성했다면, 랩을 할께.
	//그런데 30프로를 달성하지 못하면, 춤을 출께.
	percent := 0.2
	if percent >= 0.3 {
		rap()
	} else {
		dance()
	}
}
```
특이하게도 자바나 코틀린과는 달리 조건이 들어오는 부분에 `(`와 `)`이 없다.      

여러분들이 만일 플러그인을 깔았따면 괄호로 감싸도 저장시에 사라지는 마법을 볼 수 있다.      

다만 이런 방식은 불가능하다.

```go
if percent >= 0.3 
    song()
else 
    dance()
```
또한 이런 코드가 다른 언어에서는 허용되지만 `go`에서는 허용하지 않는다.

```kotlin
fun main() {

    if(a > 0)
    {
        //TODO
    }

}
```
무조건 다음과 같이

```go
// if percent >= 0.3 
// {
// 이렇게 중괄호를 들여쓰기하면 오류가 난다.
// 조건과 중괄호 {가 같은 라인에 있어야 한다.
if percent >= 0.3 {
    song()
} else {
    dance()
}
```

만일 여러개의 조건이 달린다면 어떻게 할 것인가?

```go
package main

import (
	"fmt"
)

func main() {
	flag := 10
	if flag < 5 {
		fmt.Println("5보다 작은 숫자")
    // 조건문과 중괄하 {가 같은 라인에 있어야 하는 것은
    // else if..else문에서도 적용된다.
    //  중괄호 {에 대한 들여쓰기를 허용하지 않는다.
	} else if flag == 5 {
		fmt.Println("딱 5")
	} else {
		fmt.Println("5보다 큰 수가 왔다")
	}
}
```
다른 언어와 같이 `if ~ else if ~ else`의 형식은 동일하다.   

## 쇼트서킷

이미 다른 언어를 잘 다루는 분이시라면 참 지루한 이야기지만 여러 개의 조건을 검사해야 하는 경우가 존재한다.

쇼트 서킷은 다른 언어에서도 마찬가지인데 어느 한쪽의 조건을 검사할때 조건에 맞다면 그냥 나머지를 검사하지 않는 방식을 취한다.

어째든 논리 연산자 `&&`와 `||`가 있다.      

우선 이 쇼트 서킷에 대해서 먼저 알아보자.

### &&

`&&`논리 연산자는 `AND`조건으로 양쪽 조건이 전부 `true`일때 최종적으로 참으로 간주한다.

여기서 쇼트 서킷은 양변 중 좌변이 `false`라면 이 조건은 최종적으로 결과가 `false`로 떨어지게 된다.

이것은 우변이 `true`인지 `false`인지 검사할 필요도 없다. 

물론 좌변이 `true`라면 우변도 검사할 것이다.

### ||

`||`논리 연산자는 `OR`조건으로 양쪽 조건중 하나만 `true`일때 참으로 간주한다.

여기서 쇼트 서킷은 양변 중 좌변이 `true`라면 우변이 `true`인지 `false`인지 검사하지 않는다.

마찬가지로 좌변이 `false`라면 우변의 조건을 검사할 것이다.

그렇다면 코드로 후딱 살펴보자.

```go
package main

import (
	"fmt"
)

func main() {

	age := 20

    // 좌우가 전부 참이다.
    // 이 때는 우변도 검사해서 참인지 거짓인지 검사해서 분기가 될것이다.
	if age > 10 && age < 30 {
		fmt.Println("참")
	} else {
		fmt.Println("거짓")
	}

    // 쇼트 서킷 -> &&에서 좌변이 거짓이기 떄문에 우변이 참이든 거짓이든 검사하지 않을 것이다.
    // 
	if age > 20 && age < 30 {
		fmt.Println("참")
	} else {
		fmt.Println("거짓")
	}

    // 쇼트 서킷 -> 좌변은 참이다. 따라서 우변을 검사하지 않을 것이다.
    // 우변이 거짓이라면 이때 좌변이 참인지 거짓인지 검사이후 분기가 될것이다.
    if age > 10 || age < 20 {
		fmt.Println("참")
	} else {
		fmt.Println("거짓")
	}

}
```
그렇다면 진짜 이 쇼트 서킷이 적용되는지 한번 함수를 호출하도록 만들어 보자.

```go
package main

import (
	"fmt"
)


func overTen(age int) bool {
	fmt.Println("overTen call..")
	return age > 10
}

func overTwenty(age int) bool {
	fmt.Println("overTwenty call..")
	return age > 20
}

func underTwenty(age int) bool {
	fmt.Println("underTwenty call..")
	return age < 20
}

func underThirty(age int) bool {
	fmt.Println("underThirty call..")
	return age < 30
}

func main() {

	//if age > 20 && age < 30 {
	if overTwenty(age) && underThirty(30) {
		fmt.Println("참")
	} else {
		fmt.Println("거짓")
	}

	//if age > 10 || age < 20 {
	if overTen(age) || underTwenty(age) {
		fmt.Println("참")
	} else {
		fmt.Println("거짓")
	}

}
```
실제로 이 코드를 보면 `쇼트 서킷`이 일어난다.

```
overTwenty call..
거짓
overTen call..
참
```
결과를 보면 `&&`에서 좌변의 함수가 호출되고 반환값이 `false`이기 때문에 우변의 함수는 호출되지도 않고 거짓으로 분기를 탄다.

`||`도 마찬가지로 좌변의 함수가 호출되고 반환값이 `true`가 되기 때문에 우변의 함수는 호출하지도 않고 참으로 분기를 타는 것을 볼 수 있다.

이러한 `쇼트 서킷`에 의한 특징을 잘 알아야 우리가 원하는 정확한 결과를 도출할 수 있는 로직을 작성할 수 있다.

## inline if
자바나 코틀린과는 달리 `if`문에 초기문을 넣을 수 있다.

다음과 같은 코드가 있다고 생각해 보자.
```go
package main

import (
	"fmt"
)

func main() {

	musician, isYour := choiceYours()

	if isYour {
		fmt.Println(musician)
	}

}

// 뮤지션의 이름과 해당 뮤지션을 좋아하는지에 대한 불리언 값을 반환한다.
func choiceYours() (string, bool) {
	return "ESENSE", true
}
```
여기서 `musician, isYour := choiceYours()`이 부분을 `if`조건에 인라인할 수 있다.

```go
package main

import (
	"fmt"
)

func main() {

	//musician, isYour := choiceYours()

	if musician, isYour := choiceYours(); isYour {
		fmt.Println(musician)
	}

	// 인라인한 경우에는 해당 변수가 if문의 scope로 한정되기 때문에
	// if문의 블락을 벗어나면 접근이 불가능하다.
	//fmt.Println(musician)

}

func choiceYours() (string, bool) {
	return "ESENSE", true
}
```
다음과 같이 조건 초기문을 인라인해서 사용할 수 있다.

다만 이 방식을 사용할 때는 주석에도 설명했지만 초기문에서 사용하는 변수를 `if`문의 영역밖에서는 접근이 불가능해진다.

따라서 이 변수가 `if`문 안에서만 유효하다면 인라인을 적용할 수 있고 그게 아니라며 밖으로 빼놔야 한다.

결국 이 인라인 코드는 상황에 따라 적절하게 적용하면 코드를 좀 더 깔끔하게 작성할 수 있다.

# for
프로그래밍을 하다보면 반복되는 작업을 수행해야 한다.        

`go`에서는 `for`문을 통해서 이것을 해결한다.      

다만 위에서 언급했듯이 `while`문이 없는데 이것은 `for`문의 다양한 형식을 통해서 구현할 수 있다.

그래서 어떻게 보면 다른 언어와는 다르게 `for`는 `go`에서 유일한 반복문이다.

가장 기본적인 것부터 한번 작성을 해 보자.

우선 `for`문을 작성할 때는 `if`와 마찬가지로 `(`와 `)`을 사용하지 않는다.

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop index is ", i)
	}
}
```
가장 기본적인 작성 방식은 다른 언어와 똑같다. 

다만 `(`와 `)`을 사용하지 않을 뿐이다.

여기서 변수 초기문을 생략할 수 있다.

```go
package main

import (
	"fmt"
)

func main() {
	j := 0
    
    // 변수 초기화 부분을 생략하더라도 구분자 ;를 적어준다. 
	for ; j < 10; j++ {
		fmt.Println("loop index is ", j)
	}
}
```
여기서 마지막 후처리 부분을 생략해서 작성할 수도 있다.

```go
package main

import (
	"fmt"
)

func main() {
    for k := 0; k < 10 ; {
		fmt.Println("loop index is ", k)
		k++
	}
}
```
뒤에 변수 `k`를 증감시키는 후처리 부분을 생략하고 코드 블락 안으로 변경했다.     

이때도 이 후처리 부분을 생략하더라도 구분자 `;`를 표시해줘야 한다.

당연히 이것은 앞서 배웠던 형식을 조합해서 조건문만 작성할 수도 있다.

```go
package main

import (
	"fmt"
)

func main() {
	u := 0
	for u < 10 {
	//for ; u < 10 ; {
		fmt.Println("loop index is ", u)
		u++
	}
}
```
이런 조건문만 있는 경우에는 주석 처리된 부분처럼 작성할 수 있고 구분자 `;`를 생략할 수 있다.

여러분들의 IDE에 깔린 플러그인은 저장시에는 구분자 `;`를 생략한다.     

아마도 이런 경우에는 구분자 `;`를 생략하는 코드 스타일을 추천하기 때문이 아닌가 싶다.

여기서 바로 이 조건문만 있는 형식이 `while`가 상당히 유사하다는 것을 알 수 있다.    

만일 다음과 같이 작성하게 되면 무한 로푸를 생성하게 된다.

```go
package main

import (
	"fmt"
)

func main() {
	for {
	//for true {
		fmt.Println("infinity!!")
	}
}
```
여기서 조건문 구분에 `true`를 넣었지만 생략가능하다.

## continue & break
다른 언어와 마찬가지로 반복문을 계속 실행할것인지 종료할 것인지 결정하는 키워드가 존재한다.

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("when index is 5, continue for loop")
			continue
		}

		if i == 8 {
			fmt.Println("when index is 8, exists for loop")
			break
		}
		fmt.Println("index is ", i)
	}
}
```
위 코드에서는 `index`가 5일 때는 메세지를 하나 찍는다. 

하지만 `continue`키워드를 만나면 인덱스 정보를 찍는 부분까지 가지 않고 다시 루프를 돌게 된다.

따라서 `index is 5`라는 로그가 찍히지 않게 된다.

그러다가 8일 때는 메세지를 하나 찍고 루프가 종료된다.       

## 중첩 for문내에서의 break
간혹 중첩된 `for`문을 사용할 경우가 생긴다.

여기서 단순히 루프만 도는 경우라면 상관없지만 중첩된 `for`문에 `break`키워드를 사용하는 경우에는 그 특징을 알아야 한다.

다음 코드를 보자.
```go
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("index i is ", i)
		
		for j := 0; j < 10; j++ {
			if j == 6 {
				fmt.Println("when j is 6, exists for loop")
				break
			}
			fmt.Println("index j is ", j)
		}
	}

}

```
처음 자바를 배울 때 이 코드를 이해하지 못했다.      

만일 `break`키워드를 사용하지 않는다면 총 `10 * 10`의 루프가 돈다.

중첩된 내부 `for`문내에서 `break`키워드를 사용하면 전체 루프가 종료될 것이라는 착각때문이었다.

하지만 이 루프문은 총 `10 * 7`, 즉 70번의 루프가 돈다.

그래서 대부분의 다른 언어에서도 마찬가지이지만 이런 경우에는 어떤 플래그 값을 통해서 중첩된 내부 `for`문에서 특정 조건을 만족하면 전체 루프를 종료하도록 작성해야 한다.

```go
package main

import (
	"fmt"
)

func main() {

	stop := false

	for i := 0; i < 10; i++ {
		if stop {
			fmt.Println("중첩된 내부 for문에서 조건을 만족해서 루프를 빠져나온다.")
			break
		}

		fmt.Println("index i is ", i)

		for j := 0; j < 10; j++ {
			if j == 6 {
				fmt.Println("when j is 6, exists for loop")
				stop = true
				break
			}
			fmt.Println("index j is ", j)

		}
	}

}
```
플래그인 `stop`이라는 불리언 값을 `false`로 선언하고 중첩된 내부 `for`문에서 조건에 만족하는 경우 이 플래그값을 변경하고 루프를 벗어난다.

그러면 외부 `for`문이 다시 순회할 때 이 플래그값에 의해 외부 루프도 종료하게 된다.

### label like kotlin

코틀린에서는 이런 플래그 변수를 활용하는 방법이외에 레이블을 활용해서 `break`의 영역을 어디까지 전파할지 결정하는 방법이 있다.

[관심있다면 코틀린에서 어떻게 사용하는지 보러가자](https://github.com/basquiat78/kotlin-basic-for-you/tree/main/code/controlflow#%EC%97%AC%EA%B8%B0%EC%84%9C-%EC%9E%A0%EA%B9%90)

`go`에서도 이 라벨링 기법을 사용할 수 있다.

```go
package main

import (
	"fmt"
)

func main() {
	OuterLoop:
        for i := 0; i < 10; i++ {
            fmt.Println("index i is ", i)
            for j := 0; j < 10; j++ {
                if j == 6 {
                    fmt.Println("for loop break will be propagate OuterLoop")
                    break OuterLoop
                }
                fmt.Println("index j is ", j)

            }
        }
}
```
다음과 같이 라벨링을 통해서 중첩된 내부 `for`문에서 `break`키워드를 만났을 때 어느 영역까지 전파할찌 결정한다.     

특히 지금같이 2중이 아닌 3중/4중 중첩된 루프문이라면 안쪽 깊숙한 곳에서 나를 감싸는 여러 루프문중 어디까지 종료할 지 결정하는 방법이다.

확실히 이 방법이 플래그 변수의 세팅과 체크 로직이 들어가는 방식보다는 코드 작성이 수월해 보인다.      

하지만 다른 사이트나 고랭 관련 이 부분은 오히려 플래그 변수를 활용한 방식을 권장한다.       

하지만 가장 더 좋은 방식은 이런 루프문들을 하나의 함수로 묶어서 함수를 통해서 작성하는 방식이 더 수월해 보인다.

```go
package main

import (
	"fmt"
)

func innerFor() bool {
	for i := 0; i < 10; i++ {
		if i == 6 {
			fmt.Println("외부 for문까지 루프 종료 명령이 떨어진다.")
			return false
		}
		fmt.Println("inner for index i is ", i)
	}
	return true
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("outer for index i is ", i)
		if isContinue := innerFor(); !isContinue {
			break
		}
	}
}
```
다음과 같이 코드를 분리해서 작성하는 방법이 깔끔하면서도 유지보수하기 쉬운 코드가 될 것이다.

# switch
`switch`역시 다른 언어와 비슷한 방식으로 동작한다.      

이 `switch`문은 자바보다는 코틀린과 동작 방법이 비슷한데 이유는 조건에 맞으면 해당 케이스 로직을 수행하고 끝나기 때문이다.     

일단 자바17의 경우에는 좀 다른 얘기이긴 한데 코틀린처럼 이 `switch`문을 표현식으로 사용할 때와 그렇지 않을 때 차이가 있다.

일반적인 `switch`문은 각 각의 경우에 특정 케이스만 수행할려면 여전히 `break`를 작성해 줘야 한다.

하지만 표현식으로 사용하게 되면 이 `break`가 없어도 된다.     

사실 이것은 좀 개발하는데 혼란을 줄 수 있다. 

~~차라리 그냥 `break`키워드 없이 사용하면 좋겠구먼....~~

어째든 `go`에서는 코틀린과 비슷한 방식으로 작동한다.       

하지만 이와 관련 다른 내용은 후반부에 다시 설명하기로 하고 기본적인 사용법부터 보자.

```go
package main

import (
	"fmt"
)

func main() {
	genre := "Jazz"

	switch genre {
	case "Hiphop":
		fmt.Println("비트 주세요~")
	case "Rock":
		fmt.Println("좌아아아앙~두루루 다닥 두닥 둥둥둥둥 예이에~~~~~~아아아아아!")
	case "Jazz":
		fmt.Println("Fly Me To~ the Moo~~n")
	case "Electronic":
		fmt.Println("칙칙 두두둥 칙칙 두두둥")
	default:
		fmt.Println("k-Pop")
	}
}
```
`switch`문은 `if - else if - else`가 복잡해질 때 사용하면 확실히 가독성이 좋다.      

자바에서는 `enum`과 함꼐 사용할 때 강점이 드러난다.     

아무래도 조건이 늘어날 떄 확장하기가 `if`보다는 확실히 좋기 때문이다.    

어째든 무엇을 쓰든 상황에 맞게 사용하면 될 것이다.    

특징은 다른 언어들과 비슷하기 때문에 다음과 같이 `case`부분을 묶어서 사용할 수 있다.

```go
package main

import (
	"fmt"
)

func main() {
	instrumental := "Bass"

	switch instrumental {
	case "Guitar", "Bass":
		fmt.Println("줄을 사용하는 현악기")
	case "Piano", "Fender Rhodes", "Synthesizer":
		fmt.Println("건반 계열의 악기")
	case "Drums", "Percussions":
		fmt.Println("타악기")
	default:
		fmt.Println("그외의 악기들")
	}
}
```

`case`에 조건문을 걸어서 사용하는 방법도 가능하다.

```go
package main

import (
	"fmt"
)

func main() {
    myAge := 21
	//switch true {
	switch {
	case myAge > 20 || myAge < 30:
		fmt.Println("20대시군요")
	case myAge == 21:
		fmt.Println("21살")
	case myAge > 10 && myAge < 20:
		fmt.Println("10대")
	default:
		fmt.Println("hoaxy.. 틀딱일까요??")
	}
}
```
여기서 특이점이 하나 보이는데 코드를 보면 `myAge == 21`도 조건에 맞는다.

하지만 먼저 첫 번째 `case`에서 로직이 실행되면서 종료가 되기 때문에 이 경우는 검사하지 않는다.

이 방식에서 이런 생각이 들 수 있다.

```
혹시 다른 제어문처럼 초기문 사용 쌉가능??
```

쌉가능!

```go
package main

import (
	"fmt"
)

func myAgeIs() int {
	return 21
}

func main() {
    // 아래 코드도 가능하다. 
    //switch age := myAgeIs(); true {
    // true를 생략했지만 구분자 `;`는 필요하다.
	switch age := myAgeIs(); {
	case age > 20 || age < 30:
		fmt.Println("20대시군요")
	case age == 21:
		fmt.Println("21살")
	case age > 10 && age < 20:
		fmt.Println("10대")
	default:
		fmt.Println("hoaxy.. 틀딱일까요??")
	}
}
```
물론 저 방식이 통용되는 이런 방식도 가능하다.

```go
package main

import (
	"fmt"
)

func chocieMusician() (musician string, genre string) {
	musician = "QM"
	genre = "Hiphop"
	return
}

func main() {
    switch musician, genre := chocieMusician(); genre {
	case "Hiphop":
		fmt.Println(musician, genre)
		fmt.Println("비트 주세요~")
	case "Rock":
		fmt.Println(musician, genre)
		fmt.Println("좌아아아앙~두루루 다닥 두닥 둥둥둥둥 예이에~~~~~~아아아아아!")
	case "Jazz":
		fmt.Println(musician, genre)
		fmt.Println("Fly Me To~ the Moo~~n")
	case "Electronic":
		fmt.Println(musician, genre)
		fmt.Println("칙칙 두두둥 칙칙 두두둥")
	default:
		fmt.Println(musician, genre)
		fmt.Println("k-Pop")
	}
}
```
나중에 배울테지만 자바의 `enum`스타일을 적용할 수 있다.


```go
package main

import (
	"fmt"
)

type Genre int

const (
	Hiphop Genre = iota
	Jazz
	Rock
	Electronic
	Techno
	Kpop
	Jpop
)

func main() {
    genreType := Jpop
	switch genreType {
	case Hiphop:
		fmt.Println("비트 주세요~")
	case Jazz:
		fmt.Println("Fly Me To~ the Moo~~n")
	case Rock:
		fmt.Println("비트 주세요~")
	case Kpop:
		fmt.Println("내가 만든 쿠키! 너를 위해 구웠지, but you know that it ain't for free, yeah")
	case Jpop:
		fmt.Println("기무찌!!!")
	default:
		fmt.Println("씽어쏭~~")
	}
}
```

## break & fallthrough
`switch`문에서도 자바처럼 `break`키워드가 존재한다.      

하지만 `break`키워드를 명시하든 명시하지 않든 동작은 똑같다.     

따라서 해당 키워드는 생략이 가능하다.

하지만 자바처럼 그 다음 `case`문을 순차적으로 실행할 필요성이 있을 수도 있다.

~~사실은 그런 경우를 본적이 없다~~

그럴 떄 사용할 수 있는 키워드가 `fallthrough`이다.


```go
package main

import (
	"fmt"
)

func main() {
    myChoice := "coffee"
    //그러나 나는 우유도 좋아~
	switch myChoice {
	case "coffee":
		fmt.Println("i love coffee")
		fallthrough
	case "milk":
		fmt.Println("i like milk")
	case "smoothe":
		fmt.Println("so cool~ smoothe")
	default:
		fmt.Println("nothing choice")
	}
}
```
만일 선택한 음료가 커피뿐만 아니라 우유도 좋아한다면 `fallthrough`키워드를 작성하면 된다.    

이렇게 하면 다음 `case`의 조건을 무시하고 해당 로직을 실행한다.

만일 다른 `case`에도 이 키워드를 사용하게 되면 순차적으로 다음 `case`를 실행한다.

위 코드도 마찬가지로 `i love coffee`를 찍고 다음 `i like milk`까지 찍고 `switch`문이 종료된다.

# At a Glance

`go`에서 제공하는 제어 흐름에 대한 방식을 배워보았다.     

심플하면서도 `()`와 `)` 괄호를 사용하지 않고 자바와는 달리 다양한 형식을 제공하기 때문한다.    

자바같은 언어에서 넘어 오면서 적응하기 힘들었던 형식이기도 하다.

하지만 익숙해지면 그만큼 다양한 형식의 코드 스타일을 제공하기 때문에 상황에 따른 유연함을 준다.

다음에는 배열에 대해서 알아볼 것이다.

