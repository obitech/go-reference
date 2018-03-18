package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
	// This will call the rectangle initializer even when we don't use the package
	// _ "geometry/rectangle"
)

// Blank identifier silences errors, in case we don't use the rectangle package below
// var _ = rectangle.Area

// Package variables get initialized first
var rectLen, rectWid float64 = 3, 7

// init function checkfs if len, wid > 0
func init() {
	println("geometry package initialized")

	if rectLen < 0 {
		log.Fatal("rectLen < 0 !")
	}

	if rectWid < 0 {
		log.Fatal("rectWid < 0 !")
	}
}

func number() int {
	num := 15 * 5
	return num
}

func main() {
	fmt.Println("-- Geometrical shape properties --")

	fmt.Printf("Area of rectangle: %.2f\n", rectangle.Area(rectLen, rectWid))
	fmt.Printf("Diagonal of rectnagle: %.2f\n", rectangle.Diagonal(rectLen, rectWid))

	// if statement; condition { ... }
	// num is initialized and scoped to if statement
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Standard lop
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	// While loop
	i := 0
	fmt.Println()
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println()
	// Simple switch
	count := 1

	for count <= 5 {
		switch count {
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("Index")
		case 3:
			fmt.Println("Middle")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
		default:
			fmt.Println("Not allowed")
		}
		count += 1
	}

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}
