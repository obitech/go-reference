package main

import (
    "fmt"
)

func finished() {
    fmt.Println("Finished finding largest")
}

// This will find largest number in slice of ints
func largest(nums []int) {
    // This will execute before largest function returns
    defer finished()
    fmt.Println("Starting finding largest")

    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println("Largest number in ", nums, "is", max)
    fmt.Println("End of largest func")
    // finished() will be called here because of defer
}

func printA(a int) {
    fmt.Println("Value of a IN deferred function:", a)
}

func main() {
    fmt.Println("Start of main")
    nums := []int{12, 9512, 291, 501923, 502, 3}
    fmt.Println("Calling largest func")
    largest(nums)
    fmt.Println("Back in main after largest func")
    fmt.Println("End of main")

    a := 5
    // Arguments of deferred function are evaluated when defer statement is executed
    defer printA(a)
    a = 10
    fmt.Println("Value of a BEFORE deferred functions", a)

    // defer calls are added onto a stack (LIFO)
    name := "Alex"
    fmt.Printf("Original string: %s\n", string(name))
    fmt.Printf("Reversed string: ")
    for _, v := range []rune(name) {
        defer fmt.Printf("%c", v)
    }
    fmt.Println()
}
