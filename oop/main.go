package main

import (
    "tutorial/oop/employee"
    "tutorial/oop/website"
)

func main() {
    e := employee.New("Sam", "Smith", 30, 20)
    e.LeavesRemaining()

    a1 := website.NewAuthor("Jim", "Jones", "Golang programmer")
    a2 := website.NewAuthor("Alex", "Smith", "Programmer")
    p1 := website.NewPost("Inheritance in Go", "Composition instead inheritance", a1)
    p2 := website.NewPost("Concurrency", "Go is a concurrent language", a1)
    p3 := website.NewPost("Structs instead of classes", "Go doesn't have classes", a2)

    p1.Details()

    w := website.NewWebsite(p1, p2, p3)
    w.Contents()
}
