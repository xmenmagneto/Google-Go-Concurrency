package main

import "fmt"

//  read value from right and + 1, then write to left
func f(left, right chan int) {
	left <- 1 + <-right
}


func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	//  write 1 to channel right
	go func(c chan int) {
		c <-1
	}(right)
	fmt.Println(<-leftmost)
}