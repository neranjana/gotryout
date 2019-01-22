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

// the following is also valid. The last field name will be contactInfo of type contactInfo
// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }

func main() {
	// first
	// alex := person{"Alex", "Anderson"} // this is valid but not good
	// second
	// alex := person{firstName: "Alex", lastName: "Anderson"} // better way to init struct

	//third
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 90210,
		},
	}

	// fmt.Printf("%+v", jim)
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
