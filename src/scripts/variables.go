package main

import (
	"fmt"
	"reflect"
)

var (
	name, address string
	price         float64
	surname, city, age = "Brioschi", "Cagliari", 41
)

func main() {
	mainScope := 175.5
	//unusedVar := "Useless" //if i uncomment this variable i get an error at runtime because it is unused
	name = "mauri"
	fmt.Println("Name value is ", name, "and type is", reflect.TypeOf(name))
	fmt.Println("Address value is ", address, "and type is", reflect.TypeOf(address))
	fmt.Println("price value is ", price, "and type is", reflect.TypeOf(price))
	fmt.Println("Surname value is ", surname, "and type is", reflect.TypeOf(surname))
	fmt.Println("City value is ", city, "and type is", reflect.TypeOf(city))
	fmt.Println("Age value is ", age, "and type is", reflect.TypeOf(age))
	fmt.Println("mainScope value is ", mainScope, "and type is", reflect.TypeOf(mainScope))
}
