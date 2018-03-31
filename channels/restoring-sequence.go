package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Message struct {
	str string
	wait chan bool  //  channel inside a channel, signaler
}

func main() {
	//  Get Message channel by lockstep
	c := fanIn1(boring3("Joe"), boring3("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true  //  write to inner channel to block it
		msg2.wait <- true

	}
	fmt.Println("You're boring; I'm leaving")
}


//  get whoever has something to say first
func fanIn1(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <- input1
		}
	}()
	go func() {
		for {
			c <- <- input2
		}
	}()
	return c
}

// return receive-only channel of strings
func boring3(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() { //  launch the goroutine from inside the func
		//  write into channel
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d",msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<- waitForIt
		}
	}()
	return c  //  Return the channel to the caller
}