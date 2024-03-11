// global var -> init func -> main func
package main

// 1: global var
var a int = 12

const pi = 3.14

// 2: init func will be executed before the execution of mian func.
func init() {
	println("1. This is global var:", a)
	println("2. This is init func.")
}

// 3: main func
func main() {
	println("3. This is main func.")
}
