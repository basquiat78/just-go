package main

import (
	"fmt"
)

func multiReturn() (string, int, float64, bool) {
	return "first", 2, 3.3, true
}

func multiReturn2() (first string, second int, third float64, fourth bool) {
	first = "first"
	second = 2
	third = 3.3
	fourth = true
	return
}

func main() {

	a, b, c, d := multiReturn()
	fmt.Println(a, b, c, d)

	_, intValue, _, _ := multiReturn()
	fmt.Println(intValue)

	f, s, t, _ := multiReturn2()
	fmt.Println(f, s, t)
}
