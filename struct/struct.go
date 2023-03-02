package main

import (
	"fmt"
	"unsafe"
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
	var member Member
	fmt.Println(member)

	member1 := Member{"basquiat", 25}
	fmt.Println(member1)

	member2 := Member{
		"basquiat",
		25,
	}
	fmt.Println(member2)

	member3 := Member{
		name: "basquiat",
	}
	fmt.Println(member3)

	var warrior Warrior
	fmt.Println(warrior)

	member4 := Member{"basquiat", 25}
	warrior1 := Warrior{
		memberInfo: member4,
		level:      50,
		ad:         25,
	}
	fmt.Println(warrior1)
	fmt.Println(warrior1.memberInfo.name)

	novice := Novice{"basquiat", 10}
	rogue := Rogue{novice, 30, 50, 15}
	fmt.Println(rogue)
	fmt.Println(rogue.level)
	fmt.Println(rogue.name)
	fmt.Println(rogue.Novice.level)

	var field Field
	var fieldSizeSort FieldSizeSort

	fmt.Println("field : ", unsafe.Sizeof(field))
	fmt.Println("fieldSizeSort : ", unsafe.Sizeof(fieldSizeSort))

	newNovice := new(Novice)
	newNovice.name = "basquiat"
	fmt.Println(newNovice)
	fmt.Println(*newNovice)

}
