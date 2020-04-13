package main

import (
	"fmt"
)

func main() {
	//declare a slice
	fruits := make([]string, 5, 10)
	fmt.Printf("fruits length is %d with a capacity of %d\n", len(fruits), cap(fruits))

	slice := []string{"Mauri", "Mayla", "Mia", "Fabio", "Jonata", "others", "shit"}
	fmt.Printf("slice length is %d with a capacity of %d\n", len(slice), cap(slice))
	fmt.Println("slice element 2 is", slice[2])
	fmt.Println("slice values are", slice)
	family := slice[:3]
	friends := slice[3:5] //first value of the slice(2) is included, the second value(5) is excluded
	notFamily := slice[5:]
	fmt.Println("family elements are", family)
	fmt.Println("original friends elements are", friends, "and capability is ", cap(friends))
	notFamily[1] = "work"
	fmt.Println("not part of the family are", notFamily)
	friends = append(friends, "Ale")
	fmt.Println("friends elements are", friends, "and capability is ", cap(friends))
	for _, friend := range friends {
		fmt.Println("friend is", friend)
	}
	friendsAndOthers := append(friends, notFamily...)
	fmt.Println("friends and others are", friendsAndOthers)
}
