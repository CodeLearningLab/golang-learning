package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// println will automatically add a \n at the end of line
	fmt.Print("hello again")
	// print will not
	fmt.Println("hello agin again")
	fmt.Printf("hi %s, what are you doing?", "baby")
	// printf will not also.
	info := fmt.Sprintf("hi %s, what are you doing?", "boy")
	fmt.Printf(info)
}
