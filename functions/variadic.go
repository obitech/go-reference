package main

import (
    "fmt"
)

// Variadic functions can accept variable number of arguments
func find(num int, nums ...int) (int, bool) {
    fmt.Printf("Type of nums is %T\n", nums)
    found := false

    // Nums gets converted to type slice
    for i, v := range nums {
        if v == num {
            found = true
            return i, found
        }
    }

    return -1, found

}

func main() {
    num := 105
    i, found := find(num, 90, 95, 89)
    if found {
        fmt.Printf("%v found at index %v\n", num, i)
    } else {
        fmt.Printf("%v not found.\n", num)
    }

    // Normally passing a slice to variadic function doesn't work, but there is hope
    nums := []int{90, 95, 105}
    i, found = find(num, nums...)

    switch found {
    case true:
        fmt.Printf("%v found at index %v\n", num, i)
    case false:
        fmt.Printf("%v not found.\n", num)
    }
}
