package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Anson",
		lastName:  "Felder",
		contactInfo: contactInfo{
			email:   "afelder09@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("John")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
