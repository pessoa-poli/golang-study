package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A Generator is a function that returns a Channel, such as "boring"
func boring(msg string) <-chan string { //returns a receive only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// Decouples the execution of each function, they don't need to wait on one another anymore.
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	// This is very much like having a service
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ { // Because of the synchonization nature of the channels, the messages here will be print in order. If ann is ready to send a message but joe has not sent his yet, ann will have to wait.
		fmt.Printf("You say: %v\n", <-c)
	}
	fmt.Println("You are boring, I'm leaving! >:(")
}
