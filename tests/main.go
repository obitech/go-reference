// From https://dev.to/quii/learn-go-by-writing-tests

package main

import (
    "fmt"
)

func Hello(str string) string {
    if str == "" {
        str = "World"
    }

    return fmt.Sprintf("Hello, %s", str)
}

func main() {
    fmt.Println(Hello("world"))
}
