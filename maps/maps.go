package main

import (
    "fmt"
)

// Check if two maps are the same
func mapsEqual(m1 map[string]string, m2 map[string]string) bool {
    for k, v := range m1 {
        if m2[k] != v {
            return false
        }
    }

    return true
}

func main() {
    // Zero value of map is nil...
    var personSalary map[string]int

    //... has to be allocated before assignment
    if personSalary == nil {
        fmt.Println("Allocating map")
        personSalary = make(map[string]int)
    }

    personSalary["jon"] = 12000
    personSalary["jim"] = 40000

    fmt.Println("Map contents:", personSalary)

    // Shorthand initialization
    capitals := map[string]string{
        "Germany": "Berlin",
        "France":  "Paris",
    }
    fmt.Println(capitals)

    // Access
    fmt.Println("The capital of Germany is", capitals["Germany"])

    // Check if key is in map
    _, ok := capitals["USA"]
    if !ok {
        fmt.Println("Inserting key USA")
        capitals["USA"] = "Washington"
    }
    fmt.Println("The capital of the USA is", capitals["USA"])

    // Iterating over map, order not guaranteed
    for k, v := range capitals {
        fmt.Printf("Capital of %s is %s\n", k, v)
    }

    // Deleting items
    fmt.Println("Deleting...")
    delete(capitals, "Germany")
    for k, v := range capitals {
        fmt.Printf("Capital of %s is %s\n", k, v)
    }

    // Maps are reference types
    newCapitals := capitals
    newCapitals["Italy"] = "Rome"
    fmt.Println("Iterating over map newCapitals:")
    for k, v := range newCapitals {
        fmt.Printf("Capital of %s is %s\n", k, v)
    }

    fmt.Println("Iterating over map capitals:")
    for k, v := range capitals {
        fmt.Printf("Capital of %s is %s\n", k, v)
    }

    newMap := map[string]string{
        "Poland": "Warsaw",
    }

    if equal := mapsEqual(capitals, newMap); equal {
        fmt.Println("Maps are equal")
    } else {
        fmt.Println("Maps differ")
    }

}
