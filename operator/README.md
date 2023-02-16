# 연산자

`operator`, 즉 연산자는 대부분 모든 언어에서 가장 기본적인 계산 방법을 정의한다.

여러분들이 대부분 다른 언어를 배웠을텐데 go라고 특별히 다른 건 없다.     

## 산술 연산자

말 그대로 숫자 타입의 연산을 나타내는 말이다.    

이 부분은 다른 언어와 같은데 다음 표를 통해서 확인할 수 있다.

<table>
	<head>
		<th style="text-align: center">type</dt>
		<th style="text-align: center">operator</dt>
		<th style="text-align: center">operand type</dt>
	</head>
	<tr>
		<td style="text-align: center" rowspan="5">사칙 연산</td>
		<td style="text-align: center">+</td>
		<td style="text-align: center">숫자 타입, 문자열</td>
	</tr>
	<tr>
		<td style="text-align: center">-</td>
		<td style="text-align: center">숫자 타입</td>
	</tr>
	<tr>
		<td style="text-align: center">*</td>
		<td style="text-align: center">숫자 타입</td>
	</tr>
	<tr>
		<td style="text-align: center">/</td>
		<td style="text-align: center">숫자 타입</td>
	</tr>
	<tr>
		<td style="text-align: center">%</td>
		<td style="text-align: center">정수</td>
	</tr>
	<tr>
		<td style="text-align: center" rowspan="4">비트 연산</td>
		<td style="text-align: center">&</td>
		<td style="text-align: center">정수</td>
	</tr>
	<tr>
		<td style="text-align: center">|</td>
		<td style="text-align: center">정수</td>
	</tr>
	<tr>
		<td style="text-align: center">^</td>
		<td style="text-align: center">정수</td>
	</tr>
	<tr>
		<td style="text-align: center">&^</td>
		<td style="text-align: center">정수</td>
	</tr>
	<tr>
		<td style="text-align: center" rowspan="2">시프트 연산</td>
		<td style="text-align: center"><<</td>
		<td style="text-align: center">정수 << 양의 정수</td>
	</tr>
	<tr>
		<td style="text-align: center">>></td>
		<td style="text-align: center">정수 >> 양의 정수</td>
	</tr>
</table>

## 비교 연산자

변수들을 비교해서 조건에 만족하면 true, 아니면 false반환하는 연산자 역시 다른 언어와 똑같다고 보면 된다.

<table>
	<head>
		<th style="text-align: center">operator</dt>
		<th style="text-align: center">description</dt>
	</head>
	<tr>
		<td style="text-align: center">==</td>
		<td style="text-align: center">equals</td>
	</tr>
	<tr>
		<td style="text-align: center">!=</td>
		<td style="text-align: center">not equals</td>
	</tr>
	<tr>
		<td style="text-align: center"><</td>
		<td style="text-align: center">less than</td>
	</tr>
	<tr>
		<td style="text-align: center">></td>
		<td style="text-align: center">greater than</td>
	</tr>
	<tr>
		<td style="text-align: center"><=</td>
		<td style="text-align: center">less than or equals</td>
	</tr>
	<tr>
		<td style="text-align: center">>=</td>
		<td style="text-align: center">greater than or equals</td>
	</tr>
</table>

다른 언어를 알고 있다면 너무나 잘 아는 부분일 것이다. 

다들 아시겠지만 `if`, `for`, `switch`등 분기문에서 자주 사용된다.

## 논리 연산자

역시 다른 언어와 마찬가지이다.    

피연산자들끼리의 불리언 타입의 반환값을 토대로 true, false를 반환하는 연산자이다.

다만 `!`의 경우에는 피연산자가 하나이다.


<table>
	<head>
		<th style="text-align: center">operator</dt>
		<th style="text-align: center">description</dt>
	</head>
	<tr>
		<td style="text-align: center">&&</td>
		<td style="text-align: center">AND 연산자로 양쪽의 피연산자의 결과가 모두 만족해야 한다.</td>
	</tr>
	<tr>
		<td style="text-align: center">||</td>
		<td style="text-align: center">OR 연산자로 양쪽의 피연산자중 하나라도 만족하면 된다.</td>
	</tr>
	<tr>
		<td style="text-align: center">!</td>
		<td style="text-align: center">피연산자가 true이면 false를 그리고 반대라면 true를 반환하게 된다.</td>
	</tr>
</table>

## 증감 연산자
정수 타입의 변수를 1씩 증가하거나 감소하는 구분이다.

대부분 언어에서는 `전위 증감`과 `후위 증감`을 지원하는데 go는 `후위 증감`만 가능하다.

예를 들면 자바에서는

```java
public Main {
	public void increased() {
		int index = 0;
		while(index < 10) {
			index++;
			++index;
		}
	}
}
```
같은 방식을 허용한다. 

따라서 먼저 index의 값을 증가시켜서 어떤 비지니스 로직을 처리하거나 할 수 있다.

반면에 go는
```go
a :=0

a++
a--

// 컴파일 오류로 빨겋게 밑줄 쫙
//++a
//--a
```
`전위 증감`구문을 사용하면 IDE가 친절하게 잘못되었다고 알려준다.

## 실수 연산시 주의할 점

나는 비전공 프로그래머이지만 공대 출신이라 컴퓨터의 구조를 대충은 안다.

이 부분은 이과/문과를 나누지 않더라도 아는 내용이긴 하지만 컴퓨터가 알 수 있는 것은 0과 1이다.      

이진수로 표현한다는 것을 알 수 있는데 이런 이유로 실수를 계산할 때 예기치 않는 상황을 맞닿게 된다.    

자바스크립트를 예로 들면 브라우저 개발자 도구에서 콘솔에서 바로 확인가능한데

```javascript
let a = 0.3 + 0.6
let b = 0.9
a == b
```
아시는 분들은 아시겠지만 `false`가 떨어진다.

실제로 변수 `a`를 콘솔에 찍어보면 `0.8999999999999999`가 된다.

자바에서는 실수 타입으로 `double`, `float`을 예로 들수 있는데 `float`이 `double`보다 좀 더 정밀하다 해도 이런 문제를 여러분들은 볼 수 있다.

```java
public class Main {

    public static void main(String[] args) {
        double aaa = 0.1f;
        double bbb = 0.2f;
        System.out.println(aaa + bbb);

        double aaaa = 0.3f;
        double bbbb = 0.6f;
        System.out.println(aaaa + bbbb);

        float a = 0.1f;
        float b = 0.2f;
        System.out.println(a + b);

        float aa = 0.3f;
        float bb = 0.6f;
        System.out.println(aa + bb);

    }
}
//0.30000000447034836
//0.9000000357627869
//0.3
//0.90000004
```
이미 아시는 분들도 많을 것이다.  

`float`이 조금 더 정말하긴 하지만 금융권에서는 이것을 쓰기에는 위험부담이 크다.      

그래서 자바는 `BigInteger`와 `BigDecimal`을 사용하고 다양한 옵션을 통해서 세밀하게 다룰 수 있다.

그러면 go에서도 이런 녀석이 있을까 궁금할 것이다.      

이더리움 코드를 분석하거나 또는 금융 관련 이런 계산을 다루는 분들은 아마도 `big.Float`과 `big.Int`에 대해서 알 것이다.

자바에서 이 두개를 다뤄 본 분이시라면 사용방법이 약간 비슷하다는 느낌을 받을 것이다.

한번 예제를 통해서 사용해 보자. 

`big.Float`과 `big.Int`은 "math/big"패키지에 있다.

```go
package main

import (
	"fmt"
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
}
```
이 외에도 `Nextafter()`라는 함수를 통해서 오차값을 최대한 작게 해서 그 오차 범위내에 있다면 같다고 인정하게 할 수 있다.

다음 코드를 한번 보자.

```go
package main

import (
	"fmt"
	"math"
)

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

	//Nextafter를 이용한 equalUseNextafter함수를 사용해 보자.
	fmt.Println(equalUseNextafter((floatValue1 + floatValue2), floatValue3))
}

```
위 코드를 보면 `Nextafter()`를 사용하면 값이 다르더라도 오차 범위내에 있다면 같다고 판단해서 true가 떨어지는 것을 알 수 있다.

하지만 이 방법은 사실 추천하기 힘들다.      

실제 블록체인같은 경우에는 이 작은 차이라고 하더라도 고객의 자산이기에 정확한 값을 요구하게 된다.     

그래서 자바의 경우에는 무조건 `BigDecimal`과 `BigInteger`를 사용하는 것을 권장한다.      

go에서도 이런 고객의 자산과 관련된 정밀한 계산을 요한다면 `math/big`패키지내의 `big.Float`과 `big.Int`을 사용하는 것이 좋다.

# At a Glance
대부분의 언어들이 거의 비슷하기 때문에 이 부분은 그냥 쭉 읽어보는 것만으로도 상관없을 거 같다.       

다만 증감 연산자는 `후위 증감`만 가능하다는 특징만 캐치하면 될 거 같다.

다음으로는 함수에 대해서 알아보고자 한다.