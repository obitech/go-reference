// Codewars Kata: http://www.codewars.com/kata/553e8b195b853c6db4000048/train/go

package all_unique

import (
	"fmt"
)

func HasUniqueChar(str string) bool {
	strSlice := []rune(str)
	runeMap := map[rune]int{}

	for _, r := range strSlice {
		_, ok := runeMap[r]
		if ok {
			return false
		}
		runeMap[r] += 1
	}

	return true
}
