package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	joe := boring1("Joe")  //  Function returning a channel
	ann := boring1("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're boring; I'm leaving")
}

// return receive-only channel of strings
func boring1(msg string) <-chan string {
	c := make(chan string)
	go func() { //  launch the goroutine from inside the func
		//  write into channel
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d",msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c  //  Return the channel to the caller
}