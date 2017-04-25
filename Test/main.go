package main

import (
	"./stacks"
	"fmt"
)

func main() {
	var object stacks.Stack

	for i := 1; i <= 25; i++ {
		object.Push(i)
	}

	fmt.Println(" stack capacity = ", object.Cap())
	fmt.Println(" stack length = ", object.Len())
	top, _ := object.Top()
	fmt.Println(" Return end val = ", top)
	pop, _ := object.Pop()
	fmt.Println(" Delete and return end val = ", pop)
	fmt.Println(" stack length = ", object.Len())
	var key string
	fmt.Scanln(&key)
}
