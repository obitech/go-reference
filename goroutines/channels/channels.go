package main

import (
    "fmt"
    "time"
)

func hello(done chan bool) {
    fmt.Println("Hello world goroutine started")

    for i := 1; i <= 4; i++ {
        time.Sleep(1 * time.Second)
        fmt.Printf("Sleept for %ds\n", i)
    }

    fmt.Println("Writing to channel now")
    // Writing true to channel done
    done <- true
}

// Unidirectional channel, in this case send-only
func sendData(sendch chan<- int) {
    sendch <- 10
}

// Writing 0..9 to channel, then closing it
func producer(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
}

// Receiving until channel is closed
func receiver(ch chan int) {
    for {
        v, ok := <-ch
        if ok == false {
            break
        }
        fmt.Println("Received ", v, ok)
    }
}

func main() {
    // channel c can "transport" ints from one goroutine to another
    var c chan int
    if c == nil {
        fmt.Println("Channel c is nil, defining it")
        c = make(chan int)
        fmt.Printf("Type of c is %T\n", c)
    }

    // Shorthand
    // a := make(chan int)

    // Reading from channel a
    // data := <- a
    // Read is blocked until some goroutine writes to channel

    // Wrtiting to channel a
    // a <- data
    // Write is blocked untile some gorounte reads from channel

    done := make(chan bool)
    fmt.Println("Main function demo calling goroutine")
    go hello(done)

    fmt.Println("Blocking until channel will be written")
    // Reading data from done channel, blocked until our function writes to it
    _ = <-done
    fmt.Println("Main function demo end")

    /* Bidirectional channels */
    sendch := make(chan<- int)
    go sendData(sendch)
    // Below would fail, since we can't read from that channel
    //fmt.Println(<-sendch)

    convCh := make(chan int)
    // sendData will convert the channel, inside the function it's send-only but inside main it's bidirec
    go sendData(convCh)
    fmt.Println(<-convCh)

    /* Closing channels */
    ch := make(chan int)
    go producer(ch)

    // Receiving data until channel is closed
    for v := range ch {
        fmt.Println("Received ", v)
    }
}
