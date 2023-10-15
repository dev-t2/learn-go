package magazine

import "fmt"

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Subscriber struct {
	Name        string
	Rate        int
	IsActive    bool
	HomeAddress Address
}

type Employee struct {
	Name        string
	Salary      int
	Address
}

func DefaultSubscriber(name string) *Subscriber {
	s := Subscriber{Name: name, Rate: 9900, IsActive: true}

	return &s
}

func ApplyDiscount(s *Subscriber) {
	s.Rate = 4900
}

func PrintInfo(s *Subscriber) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Monthly Rate:", s.Rate)
	fmt.Println("Active:", s.IsActive)
}