package main

import "fmt"

func main() {
	c := make(chan int)

	// This will only print 0 because main will exit after first value
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// FIX: we need to close it
		close(c)
	}()

	// FIX: Also need to range over channel
	for v := range c {
		fmt.Println(v)
	}
}
