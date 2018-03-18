package main

import "fmt"

/* -------------------------------------------- */
// Interface is a set of method signatures
type VowelsFinder interface {
    FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder implicitely
func (str MyString) FindVowels() []rune {
    var vowels []rune
    for _, rune := range str {
        if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
            vowels = append(vowels, rune)
        }
    }
    return vowels
}

/* -------------------------------------------- */

type SalaryCalculator interface {
    CalculateSalary() int
}

type Permanent struct {
    id       int
    basicPay int
    bonus    int
}

type Contract struct {
    id       int
    basicPay int
}

// Permanent type now implements SalaryCalculator interface
func (p Permanent) CalculateSalary() int {
    return p.basicPay + p.bonus
}

// Contract type now implements SalaryCalculator interface
func (c Contract) CalculateSalary() int {
    return c.basicPay
}

// To calculate total expense we iterate over a slice of all employees that
// implement the SalaryCalculator interface
func totalExpense(s []SalaryCalculator) int {
    expense := 0
    for _, v := range s {
        expense += v.CalculateSalary()
    }

    return expense
}

// Empty interface, takes any type !
func describe(i interface{}) {
    fmt.Printf("Interface has type %T, value %v\n", i, i)
}

// Type assertion: get underlying value of interface
func assert(i interface{}) {
    v, ok := i.(int)
    fmt.Println(v, ok)
}

// Type switch, this works also with custom types
func findType(i interface{}) {
    switch i.(type) {
    case string:
        fmt.Printf("I'm a string: %s\n", i.(string))
    case int:
        fmt.Printf("I'm an int: %d\n", i.(int))
    case float64:
        fmt.Printf("I'm a float64: %f\n", i.(float64))
    default:
        fmt.Printf("Unknown type\n")
    }
}

/* -------------------------------------------- */
// Pointers, Structs and Interfaces

type Info interface {
    Info()
}

type Person struct {
    name string
    age  int
}

type Address struct {
    country string
    city    string
}

// Person implements Info using a value receiver
func (p Person) Info() {
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}

// Address implements Info using pointer receiver
func (a *Address) Info() {
    fmt.Printf("Country %s, City %s\n", a.country, a.city)
}

/* -------------------------------------------- */
// Multiple Interfaces, Embedding

type SayHello interface {
    SayHello()
}

type SayBye interface {
    SayBye()
}

type SayThanks interface {
    SayThanks()
}

// Embedding interface in another interface
type SayThings interface {
    SayHello
    SayBye
    SayThanks
}

func (p Person) SayHello() {
    fmt.Printf("%s says hello!\n", p.name)
}

func (p Person) SayBye() {
    fmt.Printf("%s says bye!\n", p.name)
}

func (p Person) SayThanks() {
    fmt.Printf("%s says thanks!\n", p.name)
}

func main() {
    name := MyString("Sam Anderson")
    var v VowelsFinder
    v = name
    fmt.Printf("Vowels are %c\n", v.FindVowels())

    pemp1 := Permanent{1, 5000, 200}
    pemp2 := Permanent{2, 6000, 300}
    pemp3 := Permanent{3, 3500, 150}
    cemp1 := Contract{4, 2000}
    cemp2 := Contract{5, 2500}

    // Since both Permanent and Contract are implementing SalaryCalculator, we can append them to slice of that type
    employees := []SalaryCalculator{pemp1, pemp2, pemp3, cemp1, cemp2}
    expense := totalExpense(employees)
    fmt.Printf("Total expense per month is $%d\n", expense)

    describe(employees)

    // OK
    var i interface{} = 42
    assert(i)

    // Still OK, since we 'catch' error
    var s interface{} = "Hello World"
    assert(s)

    findType("Hello")
    findType(23)
    findType(3.14)
    findType(3i)

    /* -------------------------------------------- */
    // Pointers, Structs and Interfaces

    p1 := Person{"Sam", 25}
    p1.Info()

    p2 := Person{"James", 32}
    var d1 Info
    d1 = &p2
    d1.Info()
    // d1 is a pointer to Person
    fmt.Printf("Type of d1: %T\n", d1)
    // This works too:
    d1 = p2

    a := Address{"Germany", "Berlin"}
    fmt.Printf("Type of a: %T\n", &a)
    var d2 Info

    // This would lead to an error...
    // d2 = a
    // ... because the concrete value stored in an interface is not addressable
    // ... in the end interface values are just tuples of a value and concrete type
    // ... Calling a method on an interface value executes the method of the same name on its underlying type.

    // However, this will work
    d2 = &a
    d2.Info()
    fmt.Printf("Type of d2: %T\n", d2)

    /* -------------------------------------------- */
    // Multiple Interfaces

    // Person now implements SayHello and SayBye interfaces
    tester := Person{"Alex", 28}
    tester.SayHello()
    tester.SayBye()

    embedder := Person{"Micha", 44}
    var operations SayThings = embedder
    operations.SayHello()
    operations.SayBye()
    operations.SayThanks()

}
