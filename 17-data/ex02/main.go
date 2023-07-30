package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	li *list.List
}

func (q *Queue) Push(el interface{}) {
	q.li.PushBack(el)
}

func (q *Queue) Pop() interface{} {
	var el = q.li.Front()

	if el != nil {
		return q.li.Remove(el)
	}

	return nil
}

func CreateQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	var queue = CreateQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	var el = queue.Pop()

	for el != nil {
		fmt.Printf("%v ", el)

		el = queue.Pop()
	}
}