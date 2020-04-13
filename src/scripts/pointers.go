package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	pointer := &a
	fmt.Println("a points to", &a)
	fmt.Println("a value is", a)
	fmt.Println("pointer points to", pointer, "and type is", reflect.TypeOf(pointer))
	fmt.Println("pointer stores the value", *pointer)

	changePointer(pointer)
	fmt.Println("pointer points to", pointer)
	fmt.Println("pointer stores the value", *pointer)
}

func changePointer(pointer *int) int  {
	*pointer = 20
	return *pointer
}
