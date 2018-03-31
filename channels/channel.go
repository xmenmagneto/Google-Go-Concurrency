package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	c := make(chan string)
	go boring("boring", c)  // concurrently call boring()

	//  read from channel
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}
//  put the msg into chan c
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}