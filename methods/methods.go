package main

import (
    "fmt"
    "math"
)

type Employee struct {
    name     string
    age      int
    salary   int
    currency string
}

type Rectangle struct {
    length, width int
}

type Circle struct {
    radius float64
}

type Address struct {
    city, state string
}

// Anonymous fields
type Person struct {
    first, last string
    Address
}

/*
A method is a function with a special receiver type

func (t Type) methodName(paramlist type) {

}
*/

// displaySalary() has Employee as receiver type
func (e Employee) displaySalary() {
    fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

// Methods allows definition on different receiver types
func (r Rectangle) Area() int {
    return r.length * r.width
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

// Value receiver
func (e Employee) changeName(newName string) {
    e.name = newName
}

// Pointer receiver
// Struct will not be copied, cheaper!
func (e *Employee) changeAge(newAge int) {
    e.age = newAge
}

func (a Address) printFullAddress() {
    fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

// We can define our custom type to add methods to basic ones
type MyInt int

func (a MgyInt) add(b MgyInt) MgyInt {
    return a + b
}

func main() {
    emp1 := Employee{"Elliot Alderson", 32, 500000, "$"}
    emp1.displaySalary()

    rect := Rectangle{
        length: 10,
        width:  5,
    }
    circ := Circle{
        radius: 12,
    }

    fmt.Printf("Area of rectangle: %d\n", rect.Area())
    fmt.Printf("Area of circle: %f\n", circ.Area())

    /* Pointer vs Value receivers */
    // Value receivers don't propagate change
    fmt.Printf("Name b4 change: %s\n", emp1.name)
    emp1.changeName("Michael Andrew")
    fmt.Printf("Name after change: %s\n", emp1.name)

    // Pointer receivers actually change value of struct
    fmt.Printf("Age b4 change: %d\n", emp1.age)
    // (&emp1).changeAge(51) is not needed
    (emp1).changeAge(51)
    fmt.Printf("Age after change: %d\n", emp1.age)

    p := Person{"Elon", "Musk", Address{"LA", "California"}}

    // Methods of anonymous fields can be called as if they belong to top structure
    p.printFullAddress()

    /*
       - When FUNCTION has VALUE ARGUMENT it will accept ONLY value argument
       - When METHOD has VALUE RECEIVER it will accept value arguments AND pointer arguments
    */
    num1 := MyInt(5)
    num2 := MyInt(6)

    sum := num1.add(num2)
    fmt.Println("Sum is ", sum)

}
