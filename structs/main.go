package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// simba := person{"simba", "m"}
	// fmt.Println(shriharsha)

	// tiger := person{firstName: "tiger", lastName: "m"}
	// fmt.Println(shrihari)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	tiger := person{
		firstName: "Tiger",
		lastName:  "Xyz",
		contact: contactInfo{
			email:   "abcd@gmail.com",
			zipCode: 111111,
		},
	}

	tiger.print()

	//pass by value
	tiger.updateNamePassByVal("simba")
	tiger.print()

	//pass by ref
	tigerPointer := &tiger
	tigerPointer.updateName("Simba")
	//or tiger.updateName("blah") will also work
	tiger.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateNamePassByVal(name string) {
	p.firstName = name
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}
