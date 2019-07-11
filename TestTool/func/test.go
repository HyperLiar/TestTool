package main

import "fmt"

func main() {
	var a = []interface{}{123, "abc"}

	Print(a)

	// 因为解包后，也算是interface{}类型
	Print(a...)

	var b = []int{1,2,3}
	Print(b)
	// b解包后，并不是interface{}, 而是int
	// Print(b...)
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}