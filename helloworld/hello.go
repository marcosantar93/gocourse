package main

import "fmt"

// x:=10 can't do
var y = 11
var z int32

func main() {
	n, err := fmt.Println("hello world")
	fmt.Println(n)
	fmt.Println(err)
	z = 12
	x := 10
	fmt.Println(x + y + int(z))
}
