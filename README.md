# Just Do GO!
Gopher가 되기 위한 좌충우돌 무작정 배워보기

회사에서 go를 주언어로 사용하기 때문에 오랜만에 go를 다시 시작했다.

버전도 바뀐 부분도 있고 역시 java, kotlin을 주로 사용하다 보니 익숙치 않아서 좀 헤매인 부분이 있다.     

느꼈던 점과 차이점등을 적어가는 일종의 일기장같은 레파지토리가 될 것이다.

참고로 이 레파지토리는 프로그래밍 언어 자체를 처음 입문하시는 분들에게는 불친절할 수 있다.     

대상은 특히 나처럼 자바, 코틀린, 타입스크립트같은 다른 프로그래밍 언어를 어느정도 다뤄본 분들이다.     

# Go Version

2023년 2월 16일자로 go1.20.1버전이 릴리즈됨


# AGENDA

[변수](https://github.com/basquiat78/just-go/tree/main/variable)    

[연산자](https://github.com/basquiat78/just-go/tree/main/operator)

[함수](https://github.com/basquiat78/just-go/tree/main/functions)

[제어 흐름](https://github.com/basquiat78/just-go/tree/main/controlflow)

[배열](https://github.com/basquiat78/just-go/tree/main/array)

[구조체](https://github.com/basquiat78/just-go/tree/main/struct)


# Configuration


## Choose Your IDE

일단 결론부터 말하자면 여러분이 IntelliJ The Ultimate버전을 쓰는것이 아니라면 IntelliJ에서는 사용할 수 없다.      

이유는 언제부터인가 플러그인으로 제공했던 Go가 유료버전에서만 설치할 수 있도록 정책을 바꿨기 때문이다.

따라서 vscode나 goLand같은 IDE를 사용해야만 한다.      

### vscode

[vscode download](https://code.visualstudio.com/download)

자신의 OS에 맞춰 다운로드한 이후 설치하면 된다.        

### goLand
GoLand는 jetBrain에서 만든 Go를 위한 IDE이다.      

[goLand](https://www.jetbrains.com/go/)

다만 이녀석은 30일 trial version이다.

사실 몇 일 쓰다가 그냥 vscode로 돌아온 케이스라 30일 이후 계속 쓸 수 있는지는 잘 모른다.

어쨰든 믿고 쓰는 jetBrain에서 만든 녀석이고 go에 최적화되어 있기 때문에 좋긴 하다.     

하지만 어떤 IDE를 쓰든 그건 여러분의 선택이다.

여기서는 vscode를 기준으로 설명할 것이다.     

일단 vscode를 설치하면 플러그인을 설치해야 한다.

그중에 `Go Team at Google`에서 만들어 제공하는 `Go`와 `Go Template Support`플러그인을 설치해서 사용하고 있다.

## Installation Go

[official go homepage](https://go.dev/doc/install)

자신의 os에 맞춰서 다운로드 한다.

설치 이후에 버전을 확인해서 제대로 설치되었는지 확인한다.

2023년 2월 15일을 기준으로 현재 go 버전은 `go.1.20`이다.

```
$> go version
go version go1.20 darwin/arm64

$> which go
/usr/local/go/bin/go
```

윈도우를 사용하지 않기 때문에 다음 링크로 대체해 본다.

[윈도우에 go 설치하기](https://smg7.tistory.com/96)

맥의 경우에는 설치를 하게 되면 비슷하겠지만 일반적으로 다음 경로로 go라는 폴더가 하나 생긴다.  

```
/Users/{{your c name}}/go

e.g. /Users/basquiat/go
```
나의 맥 이름이 basquiat라 저 위치에 생성이 된다.

```
go
├── bin
└── pkg
```
이 repository는 저 위의 `go`폴더 안에 `src`폴더를 하나 만들고 그 안에 `just-go`폴더를 하나 만들어서 진행한다.    

물론 이것은 예전 방법이다.

여러분들은 꼭 이런 폴더 구조를 따를 필요는 없다.    

그 이유는 `안녕 반가워`를 찍는 시작점에서 설명한다.

# 안녕 반가워

```
go
├── bin
├── pkg
└── src
    └── just-go
```
현재 나의 디렉토리 구조는 다음과 같다.

이제부터 파일을 하나 생성해서 가장 기본이 되는 '안녕 반가워'를 찍어보자.

```go
package main

import "fmt"

func main() {
	fmt.Println("안녕 반가워")
}
```

그리고 터미널에서 다음과 같이 실행을 한다.

```
just-go> go run hello.go
안녕 반가워
```
go가 1.16버전부터 go 모듈 사용이 기본으로 되어져 있다. 

그래서 go코드는 go 모듈 아래에 있어야 하는 것이 정석이다.

실제로 해당 폴더에는 go 모듈이 없지만 위 코드는 잘 동작한다.     

하지만 다음과 같이 명령어를 하나 실행해 보자.

```
# go mod init {{module name}}
just-go> go mod init hello
```
이 명령어를 실행하면 폴더내에 `go.mod`라는 모듈 파일이 하나 생성된 것을 확인할 수 있다.

해당 파일을 열어보면

```go
module hello

go 1.20
```
위에서 처럼 설정한 모듈 이름과 go 버전이 명시되어 있다.    

# Golang의 특징

## 정적 컴파일 언어
go는 정적 컴파일 언어이다. 

일단 우리는 아주 간단한 `hello.go`를 만들어서 콘솔에 찍는 작업을 했다.    

그리고 모듈도 생성해 보았다.    

다음과 같이 이 파일을 컴파일을 해보자.

```
just-go> go build
```
명령어가 완료되면 hello라는 파일이 하나 생길 것이다. 

윈도우라면 `.exe`확장자가 붙는 파일이 생성된다고 한다.

이것을 그냥 실행해 보자.

```
just-go> ./hello
안녕 반가워
```

# Strongly Typed Language

go는 `강 타입 언어`라 해서 타입에 대해 엄격하다.     

자바에서는 `auto boxing`같은 개념이 있어서 자동 타입 변환을 해주는 경우가 많지만 go에서는 그런거 없다!

다만 이런 특징으로 인해서 타입으로 생기는 문제를 미리 방지할 수 있다는 장점이 존재한다.


# Generics

이전에는 go에서는 제네릭스가 존재하지 않았다.     

하지만 이제는 제네릭스를 지원한다.

# GC
go는 자바처럼 가비지 컬렉터가 존재한다.

# 클래스가 없다? struct(구조체)가 있다.

go는 클래스가 없다. 다만 이와 비슷한 `struct`가 존재한다.

클래스처럼 이 구조체는 메서드를 가질 수 있다.

# 상속이 없다고???

상속이 없다. 

이런 특징으로 go는 OOP가 아니라고 말하는 부분도 있는데 이것은 오해이다.

# interface

상속은 없지만 대신 인터페이스가 있다.

## Start Go

일단 우리가 만든 아주 초간단 `hello.go`를 중심으로 몇 가지 특징을 알아보자.

# 패키지 선언
자바와 비슷하게 이 패키지는 코드가 속한 위치를 알려줘야 한다.         

그 중에 우리가 먼저 알아야 하는 것은 `main`이다.

이 `main`은 아주 특별한 녀석으로 자바처럼 프로그램이 시작하는 시작점을 알려준다.       

그렇기 때문에 반드시 이 `main`패키지가 존재해야 한다.      

그리고 이 `main`패키지안에는 반드시 `main()`함수 역시 정의되어야 하는 것이다.

일단 우리는 그 전에 go의 문법적인 부분을 먼저 배워야 하기 때문에 이 모듈에 대해서는 차후에 다시 설명하고자 한다.

# 패키지 제공 (import)
자바와 비슷하지만 몇가지 차이점이 존재한다.   

우선 우리는 `hello.go`파일내에 다음과 같이 선언한 것을 볼 수 있다.

```go
import "fmt"
```
자바처럼 여러 개의 스탠다드 라이브러리나 직접 만든 모듈등을 import할 수 있다.     

하지만 `;`을 사용하지 않는다. 

```go
import "fmt"
import "math"
```

다만 go 공식 사이트에서는 다음과 같은 방식을 오히려 권장하기 때문에 위 방식보다는

```go
import (
    "fmt"
    "math"
    "error"
)
```
와 같이 사용할 것이다.

# Comment

이것은 다른 언어와 똑같이 `//`나 `/* */`을 사용한다.

자바에서는 java-doc을 만들기 위해서 `/** */`을 사용하고 `@param`, `@return`, `@throw`등을 이용한다.

하지만 go에서는 godoc이라는 툴을 사용해서 `API DOC`을 만들 수 있는데 이때 주석내용을 사용하게 된다.

따라서 대부분 go와 관련된 git의 코드들을 보면 

자바라면 다음과 같이 작성할 것이다.

```java

/**
 * 
 * @paran s -> int type
 * @return boolean
 * @throws Exception
 * @see basquiat.io.util.Test
 */
public boolean test(int s) throw Exception{
    // todo
}

```

하지만 go는 이런 방식으로 작성하면 된다.
```go

// this is test function
// int type s received, and return bool type
func test(s int) bool {
    //todo
}
```
godoc을 이용한 `API-DOC`은 차후에 한번 다뤄볼 생각이다.

# 사용하지 않는 변수, import는 에러처럼 보여진다.

go를 개발하다 보면 로직내에 사용하지 않는 변수가 있거나 임포트한 패키지중 사용하지 않는 패키지가 있다면 에러처럼 보여준다.

go입장에서는 불필요한 변수와 패키지는 메모리를 잡아먹는다고 생각하는 건지 모르겠다.      

어째든 처음 `vscode`를 설치할 때 go와 관련된 플러그인을 설치했다면 저장시 흥미로운 부분을 볼 수 있다.     

저장할 때 이런 사용하지 않는 변수나 패키지는 저장시 전부 지워버린다.

또한 코드간의 라인을 띄우기 위해 빈 라인을 여러 줄로 나누더라도 저장하는 한줄 띄기로 정렬을 해준다.

또한 다음처럼 
```
a,b,c,d -> a, b, c, d
```
자동으로 정렬을 해줄 것이다.

