package main

import (
	"fmt"
)

func main()	{
	type person struct {
		name string
		age int
	}

	Mauri := person {
		name: "Mauri",
		age: 41,
	}
	Mia := new(person) //everything is initialized with zero
	var Mayla person //everything is initialized with zero

	fmt.Println(Mauri.name, Mauri.age)
	fmt.Println(Mia.name, Mia.age)
	fmt.Println(Mayla.name, Mayla.age)
}
