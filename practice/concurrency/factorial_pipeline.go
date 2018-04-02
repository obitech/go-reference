// Challenge: Use pipeline pattern to execute 1000 factorial computations concurrently
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type factorial struct {
	n      int
	result int
}

// Calculates a the factorial of a given number (in.n) via a goroutine, and saves it as in.result
// Sending it back to a channel when done
func facConcc(in *factorial) chan *factorial {
	out := make(chan *factorial)
	go func() {
		defer close(out)
		total := 1
		for i := in.n; i > 0; i-- {
			total *= i
		}
		in.result = total
		out <- in

	}()
	return out
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	pipe := make(chan chan *factorial)
	go func() {
		defer close(pipe)
		for i := 1; i <= 1000; i++ {
			fac := factorial{
				n:      1 + rand.Intn(20-1),
				result: 0,
			}
			pipe <- facConcc(&fac)
		}
	}()

	for p := range pipe {
		res := <-p
		fmt.Printf("fac(%d) = %d\n", res.n, res.result)
	}
}
