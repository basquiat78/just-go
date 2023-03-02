# 배열
배열은 같은 타입의 정보들, 즉 요소들이 모여있는 일종의 집합체이다.

이 요소들은 그 요소들이 위치한 위치값인 소위, 인덱스를 갖게 된다.     

이러한 특징은 다른 언어와도 같은 개념을 공유한다.

먼저 이 배열을 어떻게 정의하는지 한번 살펴보자.

형식은 다음과 같다.

```go
var variableName [arraySize]int/float64/string/bool/...
```

`[]`뒤에 배열에 담을 타입을 정해서 선언한다.

이때 `[]`안에 들어가는 값은 배열의 사이즈를 결정한다.

다음 코드를 보자.
```go
package main

import (
	"fmt"
)

func main() {

	// 배열 선언 및 초기화
	var initIntArray [5]int
	fmt.Println(initIntArray)

	//
	var intArray1 = [5]int{1, 2, 3, 4, 5}
	var intArray2 = [5]int{1, 2, 3, 4}
	fmt.Println(intArray1)
	fmt.Println(intArray2)

}
// result
// [0 0 0 0 0]
// [1 2 3 4 5]
// [1 2 3 4 0]
```
일단 `int`타입의 배열을 선언하고 배열이 요소를 가질 수 있는 갯수를 5개로 정의했다.

이때 중괄호 `{`, `}`사이에 해당 타입의 값을 구분자 `,`로 정의하면 해당 값을 요소로 갖는 배열이 생긴다.

앞서 우리는 `변수`챕터에서 이미 각 타입의 기본값을 배운 적이 있다.     

따라서 `initIntArray`은 어떤 요소값도 설정하지 않았기 때문에 0으로 채워진 결과를 확인하게 된다.

다만 위 코드에서 `intArray2`처럼 사이즈를 5로 두었지만 4개의 요소만을 담은 경우에는 나머지를 해당 타입의 기본값으로 채우게 된다.

# 배열의 특정 인덱스에만 갑을 채워보자.

다음과 같이 배열을 선언해서 특정 인덱스에만 값을 채우는 방법이 존재한다.

```go
package main

import (
	"fmt"
)

func main() {
	var indexedArray = [5]int{2: 100, 4: 140}
	fmt.Println(indexedArray)
}
// [0 0 100 0 140]
```

`vararg`같은 표현식을 통해서 초기화한 요소에 맞추는 방법도 존재한다.

```go
package main

import (
	"fmt"
)

func main() {
	var likeVararg = [...]int{1, 2, 3}
	fmt.Println(likeVararg)
}
```
`...`로 설정하면 뒤에 배열의 초기 요소값을 정한 수만큼 알아서 배열의 사이즈를 생성해 준다.

# 배열 사이즈는 항상 상수?

자바에서는 이런 코드가 가능하다.

```java
public class Main {
    public static void main(String[] args) {
        int size = 5;
        String[] strArray = new String[size];
        for(String str : strArray) {
            System.out.println(str);
        }
		System.out.println(size);
		size = 10;
		System.out.println(size);
    }
}
```
size를 변수로 두고 배열 초기화에 사용이 가능하다.

하지만 `go`에서는 이것을 허용하지 않는다.        

다음 코드를 한번 살펴보자.

```go
package main

import (
	"fmt"
)

func main() {
	size := 5
	var arrays = [size]int
}
```
저게 가능할 거 같은데 변수로 선언된 값은 배열 사이즈로 정의할 수 없다.

그래서 이런 경우에는 상수로 선언해서 사용해야 한다.

```go
package main

import (
	"fmt"
)

func main() {
	const size int = 5
	var arrays = [size]int{11, 22, 33, 44, 55}
	fmt.Println(arrays)
}
```
이렇게 `const`로 선언한 상수의 경우에만 사용이 가능하다.

생각해 보면 변수의 경우에는 해당 값이 변경될 수 있다.      

물론 이것이 문제가 될 수 있을지 모르겠지만 아무튼 `go`에서는 변수를 배열의 사이즈에 사용할 수 없다.    

생각으로는 어떤 상황에서 문제를 일으킬 수 있는 변수를 사용하는 것이 버그를 발생시킬수 있다고 보는 것 같다.

# len()함수를 활용해 배열의 사이즈를 구하자.

일반적으로 자바나 코틀린의 경우에는 `.`을 이용해서 언어가 지원하는 함수를 활용해서 배열의 사이즈를 구하게 된다.

하지만 `go`에서는 지원하는 함수를 이용해서 배열의 사이즈를 구하게 된다.      

개인적으로 처음 이렇게 사용할 떄는 상당히 불편한 느낌을 지울 수가 없었다. 

```go
package main

import (
	"fmt"
)

func main() {
	var strArray [10]string
	//fmt.Println(strArray.len)이라든가
	//fmt.Println(strArray.length)라든가
	//fmt.Println(strArray.size)라든가
	//이게 더 편해보이는데??
	fmt.Println(len(strArray))
}
```

# 배열의 요소를 읽어보자.
다른 언어와 같은 형식을 지원한다.     

```go
package main

import (
	"fmt"
)

func main() {
	var strArray1 = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(strArray1[2])
	//fmt.Println(strArray1[5]) 
	//index 5 out of bounds [0:5]
}
```
물론 배열의 인덱스 크기보다 벗어난 인덱스 요소를 가져오면 `out of bounds`에러가 발생한다.

당연히 `for-loop`를 이용할 수 있는데 이것 역시 다른 언어와 똑같은 패턴이다.

```go
package main

import (
	"fmt"
)

func main() {
	var strArray1 = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(strArray1[2])
	//fmt.Println(strArray1[5])
	for i := 0; i < len(strArray1); i++ {
		fmt.Println(strArray1[i])
	}
}
```

# range키워드를 활용한 for-loop

여기서 우리는 `range`키워드를 활용해서 다음과 같이 사용도 가능하다.

```go
package main

import (
	"fmt"
)

func main() {
	var strArray1 = [5]string{"a", "b", "c", "d", "e"}

	// index활용이 필요하다면
	for index, value := range strArray1 {
		fmt.Println(index, value)
	}

	// 단지 요소값만 필요하다면
	for _, value := range strArray1 {
		fmt.Println(value)
	}
}
```
물론 이 경우에는 `index`와 그에 상응하는 요소값을 가져올 수 있는데 인덱스를 활용할 일 없다면 앞서 배운 언더스코어 '_'를 통해서 사용하지 않겠다고 명시할 수 있다.

# 대입연산자로 배열 카피하기

일단 `go`와 자바와의 차이점을 한번 살펴보자.

```java
import java.util.List;

public class Main {

    public static void main(String[] args) {
        String[] origin = new String[] {"a", "b"};
        String[] copied = origin;
        for(String str : origin) {
            System.out.println(str + ", ");
        }
        for(String str : copied) {
            System.out.println(str + ", ");
        }
        copied[0] = "A";
        for(String str : origin) {
            System.out.println(str + ", ");
        }
        for(String str : copied) {
            System.out.println(str + ", ");
        }
    }
}
```
자바를 해보신 분들이라면 위 결과는 대부분 아실 것이다.

보통 대입 연산자 `=`을 통해 새로운 변수에 다른 대입한 경우 `primitive type`이 아닌 경우에는 결과가 달라진다.     

즉 `copied`의 배열의 첫 번째 인덱스를 대문자로 변경하면 `origin`의 배열 역시 변경이 된다.     

복사는 되었지만 배열의 요소는 참조값이 복사된 것으로 실제로는 `shallow copy`, 즉 `얇은 복사`가 이뤄진 것이다.     

만일 여러분이 `얇은 복사(shallow copy)`와 `깊은 복사(deep copy)`에 대한 내용을 모른다면 구글신을 통해서 확실히 알아보고 가셔야 한다.

어째든 `go`에서는 다른 결과를 볼 수 있다.

```go
package main

import (
	"fmt"
)

func main() {
	var origin = [2]string{"a", "b"}
	var copied = origin
	fmt.Println(origin)
	fmt.Println(copied)

	copied[1] = "B"
	fmt.Println(origin)
	fmt.Println(copied)
}
```
위 코드를 실행해 보면

```
[a b]
[a b]
[a b]
[a B]
```
다음과 같이 `copied`의 첫 번째 배열의 요소만 변경된 것을 알 수 있다.

이렇게 보면 이것은 `깊은 복사(deep copy)`가 이뤄진 것이라고 볼 수 있다.      

## 값 복사와 참조 복사

사실 처음에 이 코드의 결과를 보면서 약간의 혼동이 왔었다.      

스택 영역에 있는 `primitive type`이야 원래 그렇다고 해도 이 경우에는 자바와 다른 결과를 보여줬기 때문이다.      

자바는 이 대입 연산자를 통해 객체를 복사하게 되면 해당 대상의 값이 대입되는 것이 아리나 대상의 참조값, 즉 주소 값을 대입하게 된다.

물론 이것은 자바를 공부하신 분들이라면 다들 아실 것이다.   

하지만 `go`에는 포인터가 존재한다.       

아직은 포인터를 배우지 않았지만 어째든 대입 연산자를 통해 `값 복사`를 할 것인지 `참조 복사`를 할 것인지 결정할 수 있다.

여기서 우리는 하나 확인할 수 있는 건 그냥 대입 연산자를 사용하게 되면 `값 복사`이뤄지게 때문에 다른 주소값을 갖게 되는 것을 알았다.

다음 코드를 보자.

```go
package main

import (
	"fmt"
)

func main() {
	var origin = [2]string{"a", "b"}
	var copied = origin // (1)
	fmt.Println(origin)
	fmt.Println(copied)

	copied[1] = "B"
	fmt.Println(origin)
	fmt.Println(copied)

	var copied1 = &origin // (2)
	copied1[0] = "A"
	fmt.Println(origin)
	fmt.Println(copied)
	fmt.Println(*copied1) // 역참조
}
```
(1)의 경우에는 `값 복사`가 이뤄진다.     

하지만 (2)의 경우를 살펴보자. 

차후에 배울테지만 `&`키워드를 복사할 변수앞에 붙이게 되면 이때는 `참조 복사`가 이뤄지게 된다.

따라서 이 `참조 복사`를 통해 복사한 `copied1`의 0번째 인덱스의 값을 대문자로 바꾸면 `origin`의 값까지 변경된 것을 확인할 수 있다.

자바나 자바스크립트 같은 언어에서 `go`로 넘어온 경우에는 이런 차이점을 확실히 인지하고 있어야 한다.

단지 배열뿐만이 아니다.     

예들 들면

```java
import java.util.List;

public class Main {

    public static class Test {
        private String name;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

        @Override
        public String toString() {
            return "Test{" +
                    "name='" + name + '\'' +
                    '}';
        }
    }

    public static void changeName(Test test) {
        test.setName("Basquiat");
    }

    public static void main(String[] args) {
        Test test = new Test();
        test.setName("basquiat");
        System.out.println(test);
        changeName(test);
        System.out.println(test);
    }
}
```
이런 예제코드를 짜야하나 싶을 정도의 코드이다.

결과는 `changeName`메서드를 지나는 순간 객체의 `name`이 변경된 것을 확인할 수 있다.     

자 그럼 이걸 `go`에서 비슷하게 해보자.


```go
package main

import (
	"fmt"
)

type Test struct {
	name string
}

func changeName(obj Test) {
	obj.name = "Basquiat"
}

func main() {
	test := Test{"basquiat"}
	fmt.Println(test)
	changeName(test)
	fmt.Println(test)
}
```
처음 이 코드를 짜고 테스트했을 때 자바랑 똑같을 것이라고 생각했지만 결과는

```
{basquiat}
{basquiat}
```
바뀌지 않는 것을 확인할 수 있다.

어째든 이것을 가능케 할려면 포인터를 써야 한다.

```go
package main

import (
	"fmt"
)

type Test struct {
	name string
}

// func changeName(obj Test) {
func changeName(obj *Test) { // (2)
	obj.name = "Basquiat"
}

func main() {
	test := Test{"basquiat"}
	fmt.Println(test)
	//changeName(test)
	changeName(&test)  // (1)
	fmt.Println(test)
}
```
(1)에서 처럼 `&`키워드를 사용해 메모리 주소값을 인자값으로 넘긴다.       

`changeName`함수에서 파라미터에 `*`를 붙여 해당 변수의 메모리에 접근할 수 있도록 한다.      

이렇게 하면 자바와 같은 결과를 얻을 수 있다.       

뜬금없이 배열의 대입 연산자를 이용한 복사 내용에서 여기까지 흘러들어와 버렸다.     

자바/코틀린이나 자바스크립트와는 다른 특징을 갖기 때문에 이 부분을 항상 인지하고 있어야 한다.

이부분은 차후 포인터에서 다시 한번 알아보자.

# At a Glance

`go`의 배열을 한번 알아봤다.      

위에서 말한 몇 가지 특징을 제외하면 사실 다른 언어와 크게 다를게 없다.        

차후 `slice`에서 이 배열을 동적으로 다루는 방법을 배울 것이다.        

다음 챕터는 `구조체`를 한번 다뤄볼 것이다.

