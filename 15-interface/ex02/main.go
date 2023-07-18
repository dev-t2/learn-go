package main

import "learn-go/15-interface/ex02/fedex"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	var sender = &fedex.FedexSender{}

	SendBook("TypeScript", sender)
	SendBook("Go", sender)
}