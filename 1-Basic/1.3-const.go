//const var = value
/**
const(
	a = b
	c = d
)
**/
/**
const(
	a = b
	c
)
**/
package main

import "fmt"

func main() {
	const a = true
	const (
		c = "a"
		d = 13
	)
	const (
		e = true
		f
	)
	fmt.Println(a, d, c, e, f)
}
