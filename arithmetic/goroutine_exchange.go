package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		stepA = make(chan bool)
		stepB = make(chan bool)
	)

	go func() {
		for i := 1; i <= 10; i = i + 2 {
			<-stepA
			fmt.Println("A", i)
			stepB <- true
		}
	}()

	go func() {
		for i := 2; i <= 10; i = i + 2 {
			<-stepB
			fmt.Println("B", i)
			stepA <- true
		}
	}()

	go func() {
		stepA <- true
	}()

	time.Sleep(10 * time.Second)
}
