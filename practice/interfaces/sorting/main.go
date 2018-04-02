package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	group := people{"Zeno,", "John", "Al", "Jenny"}
	fmt.Printf("Before: %v\n", group)
	sort.Sort(group)
	fmt.Printf("After: %v\n", group)

	s := []string{"Zeno,", "John", "Al", "Jenny"}
	fmt.Printf("Before: %v\n", s)
	sort.Strings(s)
	fmt.Printf("After: %v\n", s)

	n := []int{4, 26, 7, 8, 2, 8, 19, 12, 32, 3}
	fmt.Printf("Before: %v\n", n)
	sort.Ints(n)
	fmt.Printf("After: %v\n", n)

}
