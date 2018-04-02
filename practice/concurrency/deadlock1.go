package main

import "fmt"

func main() {
	c := make(chan int)

	// This will deadlock without a goroutine; it will block bc there is nothing to receive from channel
	// c <- 1

	// This goroutine will overstep the block, advancing in the control flow and letting Println receive
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
