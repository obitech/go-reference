// Challenge: use goroutines to calculate factorial
package main

import "fmt"

func facNorm(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func facConc(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func main() {
	n := 25
	fmt.Printf("facNorm(%d) = %d\n", n, facNorm(n))

	c := facConc(n)
	for n := range c {
		fmt.Print(n)
	}
}
