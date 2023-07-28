package post

import "fmt"

type PostSender struct{}

func (p *PostSender) Send(parcel string) {
	fmt.Printf("Post Sends %v Parcel\n", parcel)
}