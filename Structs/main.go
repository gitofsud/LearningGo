package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	// I
	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
	}
	fmt.Println(alex)

	// II
	var bob person
	fmt.Printf("%+v", bob)
	fmt.Println()

	bob.firstName = "Bob"
	bob.lastName = "Builder"
	fmt.Printf("%+v", bob)
}