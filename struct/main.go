package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	phone     int
	contact   contactInfo
	// we could also initialise as below like objects in js is key value name is same
	// contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	// declaring a person various ways

	// when we declare and intialise using this approach, we have to give all field values
	// cause for this approach compiler relies on the correct order of args
	alex := person{
		"Alex",
		"Anderson",
		123,
		contactInfo{ // note here that we have to annotate the type
			"a@b.com",
			123,
		}, // when declaring multi line structs, every line must have a comma
	}

	// depends on order of args // not recommended
	fmt.Println(alex)

	fmt.Println(alex.contact.email) // reading individual values

	// another way, more safe // we declare the key names as well, so order is not required
	// also since we use key names, we can omit some keys and those will be assgned Zero values
	jim := person{firstName: "Jim", lastName: "Halpert"}
	// here phone, contactInfo are assigned zero values
	fmt.Println(jim)

	// another way
	var manish person
	fmt.Println(manish) // this print { 0} // firstName "", lastName "", phone 0, zero values

	// we can log out individual values using Printf and "%+v"
	fmt.Printf("%+v", manish)

	// %v will also print out but only the values
	// %+v will print out detailed name-value pairs

	// print using the receiver function
	manish.print()

	refToManish := &manish
	fmt.Println("*********")

	fmt.Println("ref ---> ", refToManish)

	refToManish.updateFirstName("Manish")

	// manish.updateFirstName("Manish")
	manish.print()

	// since getting a reference to a type to call
	// with a function which works with pointer to that type is such a common use case that
	// go allows us a shortcut to call such functions directly with the type
	// instead of getting a refernce and compiler figures it out correctly

	// so we can also do
	manish.updateFirstName("Raja Bhaiya")

	manish.print()

}

// function with receiver of type struct
func (p person) print() {
	fmt.Printf("%+v", p)
}

// func with receiver to update name
// GO is a pass by value language

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}
