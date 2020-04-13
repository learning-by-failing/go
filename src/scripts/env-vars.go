package main

import (
	"fmt"
	"os" //operating system
)

func main()  {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	name := os.Getenv("LOGNAME")

	fmt.Println("Logged user is", name)
}