package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	// declaring a person various ways

	alex := person{"Alex", "Anderson"} // depends on order of args // not recommended

	fmt.Println(alex)

	// another way, more safe

	jim := person{firstName: "Jim", lastName: "Halpert"}

	fmt.Println(jim)

}
