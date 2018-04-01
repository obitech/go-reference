package main

import (
	"fmt"
)

func findSum(arr []int, sum int) bool {
	lenArr := len(arr)

	for i := 0; i < lenArr; i++ {
		for j := i + 1; j < lenArr; j++ {
			if arr[i]+arr[j] == sum {
				found := arr[i] + arr[j]
				fmt.Printf("Found pair at %d + %d = %d, Indexes: %d, %d\n",
					arr[i], arr[j], found, i, j)
				return true
			}
		}
	}

	return false
}

func main() {
	arr := []int{8, 7, 2, 5, 3, 1}
	_ = findSum(arr, 10)
}
