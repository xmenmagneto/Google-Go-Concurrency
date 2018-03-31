package main

import (
	"fmt"
	"time"
	"math/rand"
)

func randomSleep(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	randomSleep("boring")
}
