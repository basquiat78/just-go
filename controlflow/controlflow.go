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

func choiceYours() (string, bool) {
	return "ESENSE", true
}

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

func myAgeIs() int {
	return 21
}

func chocieMusician() (musician string, genre string) {
	musician = "QM"
	genre = "Hiphop"
	return
}

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
	//내가 달성치를 30프로이상 달성했다면, 랩을 할께.
	//그런데 30프로를 달성하지 못하면, 춤을 출께.
	percent := 0.3
	if percent >= 0.3 {
		rap()
	} else {
		dance()
	}

	flag := 10
	if flag < 5 {
		fmt.Println("5보다 작은 숫자")
	} else if flag == 5 {
		fmt.Println("딱 5")
	} else {
		fmt.Println("5보다 큰 수가 왔다")
	}

	age := 20
	/*
		if age > 10 && age < 30 {
			fmt.Println("참")
		} else {
			fmt.Println("거짓")
		}

		if age > 20 && age < 30 {
			fmt.Println("참")
		} else {
			fmt.Println("거짓")
		}

		if age > 10 || age < 20 {
			fmt.Println("참")
		} else {
			fmt.Println("거짓")
		}

		if age > 10 || age < 20 {
			fmt.Println("참")
		} else {
			fmt.Println("거짓")
		}
	*/

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

	//musician, isYour := choiceYours()

	if musician, isYour := choiceYours(); isYour {
		fmt.Println(musician)
	}

	// 인라인한 경우에는 해당 변수가 if문의 scope로 한정되기 때문에
	// if문의 블락을 벗어나면 접근이 불가능하다.
	//fmt.Println(musician)

	for i := 0; i < 10; i++ {
		fmt.Println("loop index is ", i)
	}

	j := 0
	for ; j < 10; j++ {
		fmt.Println("loop index is ", j)
	}

	for k := 0; k < 10; {
		fmt.Println("loop index is ", k)
		k++
	}

	u := 0
	for u < 10 {
		//for ; u < 10 ; {
		fmt.Println("loop index is ", u)
		u++
	}

	/*
		for {
			//for true {
			fmt.Println("infinity!!")
		}
	*/

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

	for i := 0; i < 10; i++ {
		fmt.Println("outer for index i is ", i)
		if isContinue := innerFor(); !isContinue {
			break
		}
	}

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

	//switch age := myAgeIs(); true{
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

	switch musician, genre := chocieMusician(); genre {
	case "Hiphop":
		fmt.Println(musician, genre)
		fmt.Println("비트 주세요~")
	case "Rock":
		fmt.Println(musician, genre)
		fmt.Println("비트 주세요~")
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

	myChoice := "coffee"
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
