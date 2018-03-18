package main

/**
 * - Goroutines are functions/methods running concurrently with other functions/methods
 * - Like a lightweight thread (only few kb per goroutine in stack size)
 * - Multiplexed to number of OS threads: 1 thread in a program with 1000s goroutines
 * - Communication via channels, preventing race conditions
**/

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Hello world goroutine")
}

func numbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}

func alphabets() {
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c", i)
    }
    fmt.Println()
}

func main() {
    /*
       // This starts a gouroutine
       go hello()

       // Running the program will only output below... why?
       // - started goroutine returns immediately and doesn't way for return values
       // - If main Goroutine terminates, program will be terminated

       fmt.Println("Main func")

       // This will fix it
       go hello()

       // Main goroutine sleeps for 1s, dirty hack!
       time.Sleep(1 * time.Second)
       fmt.Println("Main func after wait")
    */

    fmt.Println("Main started")
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("Main terminated")
}
