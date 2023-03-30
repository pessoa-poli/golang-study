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
	fmt.Println("Hi")
	c := boring("boring!") // Channels are first class values just like strings or integers
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %v\n", <-c)
	}
	fmt.Println("You are boring, I'm leaving! >:(")
}
