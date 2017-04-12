package main

import(
	"fmt"
	"./stack"
)

func main() {

	var name stack.Stack
	name.Push(35)
	name.Push(255)
	n := name.Len()
	fmt.Println(n)
}
