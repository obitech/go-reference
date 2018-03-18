// Every GO file needs to start with package
package main

import "fmt"

func print_age(age int) {
    fmt.Println("my age is ", age)
    return
}

func run_variables() {
    // Simple variable declaration
    var age int
    fmt.Println("my age is ", age)

    // Variable assignment
    age = 29
    print_age(age)

    // Initial value assignment
    var alter = 44
    print_age(alter)

    // Type inference, in this case `int` is inferred
    var anos = 55
    print_age(anos)

    // Multiple assignments
    //var width, height int = 100, 50
    //fmt.Println("width is ", width, ", height is ", height)

    // Type can be omitted
    //var width, height = 100, 50
    //fmt.Println("width is ", width, ", height is ", height)

    // Declaring variables of different type in same statement
    var (
        name   = "pedro"
        agee   = 44
        height int
    )
    fmt.Println("My name is", name, ", age is,", agee, "and height is", height)

    // Shorthand variable declaration
    test := "foo bar"
    fmt.Println("It's just", test)

    // Constants
    const alpha = "alpha"
    const get = 50
    fmt.Println(alpha, get)
}

func run_strings() {
    first := "Peter"
    last := "Lustig"
    name := first + " " + last
    fmt.Println("My name is ", name)

}

func calc_bill(price, nr int) int {
    return price * nr
}

func rectProps(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = (length + width) * 2
    return
    // area & perimeter are automatically returned when return appears
}

// Should always reside in main package
func main() {
    //run_variables()
    //run_strings()

    //var area, _ = rectProps(5, 3)
    //fmt.Println(area)
}
