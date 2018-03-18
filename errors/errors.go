package main

import (
    "fmt"
    "os"
)

/**
 * Error is an interface looking like this:
 * type error interface {
 *      Error(string)
 * }
 */

func main() {
    f, err := os.Open("/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}
