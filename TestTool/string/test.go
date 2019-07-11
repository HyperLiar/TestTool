package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "sinaweibo"

	for _, c := range []byte(s) {
		fmt.Println(c)
		fmt.Println(reflect.TypeOf(c))
	}
}
