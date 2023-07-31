package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	li *list.List
}

func CreateStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(el interface{}) {
	s.li.PushBack(el)
}

func (s *Stack) Pop() interface{} {
	var back = s.li.Back()

	if back != nil {
		return s.li.Remove(back)
	}

	return nil
}

func main() {
	var stack = CreateStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	var el = stack.Pop()

	for el != nil {
		fmt.Printf("%v ", el)

		el = stack.Pop()
	}
}