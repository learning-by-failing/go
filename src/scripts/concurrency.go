package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

func main()  {
	runtime.GOMAXPROCS(3) // create 2 process maximum
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func () {
		defer waitGrp.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Mauri")
	}()

	go func ()	{
		defer waitGrp.Done()
		fmt.Println("Ciao")
	}()

	waitGrp.Wait()
}