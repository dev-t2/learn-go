package main

import (
	"container/list"
	"fmt"
)

func main() {
	var li = list.New()

	var el1 = li.PushFront(1) 
	var el4 = li.PushBack(4)

	li.InsertAfter(2, el1)
	li.InsertBefore(3, el4)

	for el := li.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value, " ")
	}

	fmt.Println()

	for el := li.Back(); el != nil; el = el.Prev() {
		fmt.Print(el.Value, " ")
	}
}