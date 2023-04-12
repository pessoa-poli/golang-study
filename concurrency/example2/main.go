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

func main() {
	// This is very much like having a service
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ { // Because of the synchonization nature of the channels, the messages here will be print in order. If ann is ready to send a message but joe has not sent his yet, ann will have to wait.
		fmt.Printf("You say: %v\n", <-joe)
		fmt.Printf("You say: %v\n", <-ann)
	}
	fmt.Println("You are boring, I'm leaving! >:(")
}
