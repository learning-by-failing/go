package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		message := "message " + strconv.Itoa(i)
		delay := time.Second * time.Duration(rand.Intn(10))
		writeMessageNorConcurrent(message, delay)
	}
}

func writeMessageNorConcurrent(message string, delay time.Duration) {
		fmt.Println(message, "in", delay)
		time.Sleep(delay)
		fmt.Println(message, "out")
}
