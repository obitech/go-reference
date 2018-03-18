package main

import (
    "fmt"
)

// Use slices to modify arrays, not pointers to arrays!
func modify(sls []int) {
    fmt.Println("Modifying...")
    sls[0] = 99
}

func main() {
    b := 255

    // a is a pointer to int, pointing to memory address of b
    var a *int = &b
    fmt.Printf("type(a) = %T\n", a)
    fmt.Println("Address of b is ", a)

    // zero value of pointer is nil
    c := 25
    var d *int

    if d == nil {
        fmt.Println("d is ", d)

        // d pointing to c now
        d = &c

        fmt.Println("d is pointing to", d)
        fmt.Println("memory addres of c is", &c)

        // d is pointing to c, so if we dereference d, we will get the value of c
        fmt.Println("c = ", *d)

        // We can change the value of c while dereferencing d
        *d = 42
        fmt.Println("c = ", c)

        f_arr := [...]int{1, 2, 3, 24, 42}
        f_sls := f_arr[:]
        fmt.Printf("f_arr: %d\nf_sls: %d\n", f_arr, f_sls)
        modify(f_sls)
        fmt.Printf("f_arr: %d\nf_sls: %d\n", f_arr, f_sls)

        // Thank god there is no pointer arithmetic in Go !
    }
}
