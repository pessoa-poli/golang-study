package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A Generator is a function that returns a Channel, such as "boring"
func boring(msg string) <-chan Message { //returns a receive only channel of strings
	waitForIt := make(chan bool)
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

//Restoring sequence:
// Send a channel on a channel, making a goroutine wait it'd turn.
// Receive all messages then enable them again by sending on a private channel
type Message struct { // Message type that contains a channel for the reply
	str  string
	wait chan bool
}

// Decouples the execution of each function, they don't need to wait on one another anymore.
func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
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
	for i := 0; i < 5; i++ { // Because of the synchonization nature of the channels, the messages here will be print in order. If ann is ready to send a message but joe has not sent his yet, ann will have to wait.
		msg1 := <-c
		msg2 := <-c
		fmt.Printf("You say: %v\n", msg1)
		fmt.Printf("You say: %v\n", msg2)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You are boring, I'm leaving! >:(")
}
