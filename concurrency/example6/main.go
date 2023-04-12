package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SELECT: The reason channels and goroutines are built into the

// Select:
// The select statement provides another way to handle multiple channels.
// It's like a switch, but each case is a communication:
// - All channels are evaluated.
// Selection blocks until one communication can proceed, which then does.
// - If multiple can proceed, select chooses pseudo-randomly.
// A default clause, if present, executes immediately if no channel is ready.

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

// This fanIn function behaves exactly like example5, except we launch just one goRoutine
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
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
