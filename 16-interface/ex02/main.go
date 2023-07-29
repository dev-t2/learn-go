package main

import (
	"fmt"
	"learn-go/16-interface/ex02/fedex"
	"learn-go/16-interface/ex02/post"
)

type Sender interface {
	Send(parser string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var postSender = &post.PostSender{}

	SendBook("JavaScript", postSender)
	SendBook("NodeJS", postSender)

	fmt.Println()

	var fedexSender = &fedex.FedexSender{}

	SendBook("TypeScript", fedexSender)
	SendBook("Go", fedexSender)
}