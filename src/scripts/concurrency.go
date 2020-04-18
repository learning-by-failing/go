package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const messageNumber = 10

var (
	waitGrp sync.WaitGroup
)

func main() {
	waitGrp.Add(messageNumber)
	for i := 0; i < messageNumber; i++ {
		message := "message " + strconv.Itoa(i)
		delay := time.Second * time.Duration(rand.Intn(messageNumber))
		go writeMessage(message, delay)
	}
	waitGrp.Wait()
}

func writeMessage(message string, delay time.Duration) {
	go func() {
		defer waitGrp.Done()
		fmt.Println(message, "in", delay)
		time.Sleep(delay)
		fmt.Println(message, "out")
	}()
}
