package stacks

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can`t Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	temp := *stack
	if len(temp) == 0 {
		return nil, errors.New("can`t Pop() an empty stack")
	}

	x := temp[len(temp)-1]
	*stack = temp[:len(temp)-1]
	return x, nil
}
