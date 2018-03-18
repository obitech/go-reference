package main

import (
    "fmt"
)

func main() {
    // Int array with 3 length, initialized to 0
    var a [3]int

    fmt.Println(a)

    // Assign values
    a[0] = 42
    a[1] = 23
    a[2] = 2018
    fmt.Println(a)

    // Short hand declaration
    b := [3]int{3, 8, 9}
    fmt.Println(b)

    // Compile will substitute correct length
    c := [...]int{10, 11, 12, 15, 18, 23}
    fmt.Println(c)

    // c has now a custom type: [3]int !!
    fmt.Printf("Type of c: %T\n", c)

    // range returns both index and value of array
    for i, v := range c {
        fmt.Printf("index: %[1]d, value: %[2]d\n", i, v)
    }

    // Arrays are passed by value, not reference!
    fmt.Printf("a: %v\nb: %v\n", a, b)
    fmt.Println("a = b")
    a = b
    fmt.Println("b[1] = -99")
    b[1] = -99
    fmt.Printf("a: %v\nb: %v\n", a, b)

    // Create multi-dimensional arrays and assign values
    d := [3][3]int{}
    for i, v := range d {
        for j := range v {
            d[i][j] = i + j
        }
    }
    fmt.Println(d)

    // Slices wrap around arrays, they reference arrays and don't own data
    e := [5]int{}
    for i := range e {
        e[i] = i
    }
    fmt.Printf("Array: %[1]d, type: %T\n", e, e)

    // Create slice from a[1] to a[3]
    var g []int = e[1:4]
    fmt.Printf("Slice: %v, type: %T\n", g, g)

    // Creates array and returns slice of it
    f := []int{6, 7, 8}
    fmt.Printf("Slice: %v, type: %T\n", f, f)

    // Changes in slices will be reflected in underlying array
    sarr := [...]int{1, 2, 3, 4, 5}
    fmt.Println("Array sarr:", sarr)

    sslice := sarr[0:3]
    fmt.Println("sslice[0:3] of sarr:", sslice)

    fmt.Println("Changing sslice...")
    for i := range sslice {
        sslice[i] = -99
    }

    fmt.Println("sslice[0:3] of sarr:", sslice)
    fmt.Println("Array sarr:", sarr)

    // Length of slice = # of elements in slice
    // Capacity of slice = # of elements in array starting from first slice index
    h := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    h_s := h[1:3]
    fmt.Printf("Array: %v, Slice: %v\n", h, h_s)
    // Len = 2 because 2 elements
    // Cap = 8 because first element is index 1 !
    fmt.Printf("Len of slice: %v, Cap of slice: %v\n", len(h_s), cap(h_s))
    // Slices can be reslized up to capacity
    fmt.Println("Reslicing...")
    h_s = h_s[1:8]
    fmt.Printf("Slice: %v, Len of slice: %v, Cap of slice: %v\n", h_s, len(h_s), cap(h_s))

    // Creating a slice using funct make([]T, len, cap) []T
    i := make([]int, 5, 5)
    fmt.Println(i)

    // Appending to a slice
    // New array will be created, things will get copied
    cars := []string{"Audi", "VW", "Porsche"}
    fmt.Printf("cars: %v, len: %v, cap: %v\n", cars, len(cars), cap(cars))
    cars = append(cars, "Toyota", "Honda")
    // Capacity gets doubled!
    fmt.Printf("cars: %v, len: %v, cap: %v\n", cars, len(cars), cap(cars))

    /*
       Internal representation of a slice:
       type slice struct {
           Length        int
           Capacity      int
           ZerothElement *byte
       }
       --> Thus, passing it around will change the underlying array!
    */

    // As long as a slice references an array it can't be GCed !
    // Use copy(dst, src[]T) int to make a copy & for better memory management:
    // 1. Create a (big) array
    numbers := [1000]int{}
    for i := range numbers {
        numbers[i] = i
    }
    // 2. Create slice of array
    neededNumbers := numbers[:len(numbers)/2]
    // 3. numbers array can't get GCed because neededNumbers is referencing it
    // Let's create a copy so GC can clean numbers up
    numbersCpy := make([]int, len(neededNumbers))
    copy(numbersCpy, neededNumbers)
    fmt.Printf("%T\n", numbersCpy)
}
