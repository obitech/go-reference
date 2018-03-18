package main

import (
    "fmt"
    "sync"
    "time"
)

/* -- Wait Groups -- */
// Used to wait for a collection of Goroutines to finish
// Control is blocked untiall all Goroutines finish executing

func process(i int, wg *sync.WaitGroup) {
    fmt.Println("started Gouroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)

    // Decrementing the counter of the WaitGroup
    wg.Done()
}

func main() {
    no := 3

    // WwaitGroup is a struct type
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {

        // Works like a counter, here we add one to it
        wg.Add(1)

        // We need to pass the address or each goroutine will have its own WaitGroup
        go process(i, &wg)
    }

    // Blocked until counter decrements to 0
    wg.Wait()
    fmt.Println("All Goroutines finished executing.")
}
