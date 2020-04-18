package main

import (
	"fmt"
	"time"
)

/**
communication example inside the channel
fibonacci starts, then it waits for the messages from the other go routine
using the select and handling it with the case that mach (always c <- x) because the channle "quit"
is filled only at the end of the first go routine
*/
func fibonacci(c, quit chan int) {
	fmt.Println("fibonacci")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default: //the default could be simply avoid, it will loop forever -> here just a demo
			time.Sleep(1 * time.Second)
			fmt.Println("in loop")
		}
	}
	fmt.Println("fibonacci ends")
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hi", <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
