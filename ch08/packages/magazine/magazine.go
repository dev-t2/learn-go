package magazine

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
	HomeAddress Address
}