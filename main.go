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
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	//var alex person

	//	alex.firstName = "Alex"
	//	alex.lastName = "Andreson"

	//	fmt.Println(alex)
	//	fmt.Printf("%+v", alex)

	subi := person{
		firstName: "subi",
		lastName:  "bhardwaj",
		contact: contactInfo{
			email:   "subibhardwaj24@gmail.com",
			zipCode: 56789,
		},
	}

	//fmt.Printf("%+v",subi)
	subiPointer := &subi
	subiPointer.updateName("jimmy")
	subi.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
