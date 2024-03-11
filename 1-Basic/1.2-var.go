// var varName type = value
// var varNameA, varNameB string = valueA, valueB
// varName := value
// var varNameA, varNameB string := valueA, valueB
// assignment:    var = value
/**
var (
	user = "a"
	name = "aaa"
)
**/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "aaa"
	var age int = 15
	var isBoy bool = true
	var hello, init string = "world", "code"
	e, d := "c", true
	var (
		a = "b"
		c = 14
	)
	fmt.Println(reflect.TypeOf(name), name)
	fmt.Println(reflect.TypeOf(age), age)
	fmt.Println(reflect.TypeOf(isBoy), isBoy)
	fmt.Println(hello, init)
	fmt.Println(a, c)
	fmt.Println(d, e)
}
