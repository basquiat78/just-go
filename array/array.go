package main

import (
	"fmt"
)

type Test struct {
	name string
}

// func changeName(obj Test) {
func changeName(obj *Test) {
	obj.name = "Basquiat"
}

func main() {

	var initIntArray [5]int
	fmt.Println(initIntArray)

	var intArray1 = [5]int{1, 2, 3, 4, 5}
	var intArray2 = [5]int{1, 2, 3, 4}
	fmt.Println(intArray1)
	fmt.Println(intArray2)

	var indexedArray = [5]int{2: 100, 4: 140}
	fmt.Println(indexedArray)

	var likeVararg = [...]int{1, 2, 3}
	fmt.Println(likeVararg)

	// 이거 안된다.
	//size := 5
	//var arrays = [size]int{1, 2, 3, 4, 5}
	const size int = 5
	var arrays = [size]int{11, 22, 33, 44, 55}
	fmt.Println(arrays)

	var strArray [10]string
	fmt.Println(len(strArray))

	var strArray1 = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(strArray1[2])
	//fmt.Println(strArray1[5])
	for i := 0; i < len(strArray1); i++ {
		fmt.Println(strArray1[i])
	}

	for index, value := range strArray1 {
		fmt.Println(index, value)
	}

	for _, value := range strArray1 {
		fmt.Println(value)
	}

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
	fmt.Println(*copied1)

	test := Test{"basquiat"}
	fmt.Println(test)
	//changeName(test)
	changeName(&test)
	fmt.Println(test)
}
