package main

import (
    "fmt"
    "tutorial/structures/computer"
)

// Structs is a custom type, representing collection of fields
type Employee struct {
    firstName, lastName string
    age, salary         int
}

// Anonymous fields
type Person struct {
    string
    int
}

// Nested structs
type Company struct {
    name string
    ceo  Person
}

type Address struct {
    city, state string
}

// Promoted fields
type House struct {
    owner string
    Address
}

type Name struct {
    first, last string
}

func equalNames(s1 Name, s2 Name) bool {
    if s1 == s2 {
        fmt.Printf("%v and %v are equal\n", s1, s2)
        return true
    } else {
        fmt.Printf("%v and %v differ\n", s1, s2)
        return false
    }
}

func main() {
    // Named structs
    emp1 := Employee{
        firstName: "Sam",
        age:       25,
        salary:    500,
        lastName:  "Anderson",
    }
    emp2 := Employee{"Thomas", "Peterson", 29, 800}
    fmt.Println("Employee 1:", emp1)
    fmt.Println("Employee 2:", emp2)

    // Anonymous structs
    emp3 := struct {
        firstName, lastName string
        age, salary         int
    }{"Andrea", "Nikita", 30, 900}
    fmt.Println("Employee 3:", emp3)

    // Empty structs are assigned zero values of fields
    var emp4 Employee
    fmt.Println("Employee 4:", emp4)

    // Only defining specific fields
    emp5 := Employee{
        firstName: "John",
        age:       30,
    }
    fmt.Println("Employee 5", emp5)

    // Accessing/Defining specific fields
    emp5.lastName = "Wayne"
    fmt.Println("emp5 salary: ", emp5.salary)
    fmt.Println("emp5: ", emp5)

    // Pointers to structs
    emp6 := &Employee{"Mona", "Lisa", 800, 600000}
    fmt.Printf("Type of emp6: %T\n", emp6)
    // Those two are the same
    fmt.Println("First name:", (*emp6).firstName)
    fmt.Println("First name:", emp6.firstName)

    p1 := Person{"Rashad", 32}
    fmt.Println(p1)

    // By default, name of anonymous field equals type
    fmt.Println(p1.string)

    // Testing nested structs
    c1 := Company{"Best Company", p1}
    fmt.Println(c1)
    c1.ceo.string = "Joseph"
    fmt.Println(c1)

    // House has an promoted field Address
    var h House
    h.owner = "Rashad"
    // Defining the promoted field
    h.Address = Address{
        city:  "Munich",
        state: "Bavaria",
    }
    fmt.Println("House h:", h)
    // We can now access the Address fields as if they were part of House
    fmt.Println("City:", h.city)
    fmt.Println("State:", h.state)

    // Testing exported structs
    var spec computer.Spec
    spec.Maker = "Apple"
    spec.Price = 13000
    fmt.Println("Spec:", spec)

    // Structs are comparable if each of their fields are comparable
    name1 := Name{"Steve", "Jobs"}
    name2 := Name{"Steve", "Jobs"}
    name3 := Name{"Don", "Corleone"}
    equalNames(name1, name2)
    equalNames(name1, name3)
}
