package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

type pair struct {
	Key   string
	Value int
}

type pairList []pair

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Less(i, j int) bool {
	return p[i].Value < p[j].Value
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// From https://stackoverflow.com/a/18695740
func sortMapByValue(m map[string]int) pairList {
	// Create a list where store our K/V pairs
	out := make(pairList, len(m))
	var i int

	// Insert each pair
	for k, v := range m {
		out[i] = pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(out))
	return out
}

func sortMapByKeys(m map[string]int) pairList {
	store := make([]string, len(m))
	var i int

	// Store keys in list
	for k := range m {
		store[i] = k
		i++
	}
	sort.Strings(store)

	// Enter sorted strings in pairList
	out := make(pairList, len(m))
	for i, v := range store {
		out[i] = pair{v, m[v]}
	}
	return out
}

func main() {
	books := map[string]string{
		"moby":       "https://www.gutenberg.org/files/2701/2701-0.txt",
		"words":      "https://raw.githubusercontent.com/dwyl/english-words/master/words.txt",
		"words_file": "./words.txt",
		"test_file":  "./test.txt",
		"moby_file":  "./2701-0.txt",
	}

	// Open book
	file, err := ioutil.ReadFile(books["moby_file"])
	check(err)
	str := string(file)
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	wordCount := make(map[string]int)

	for scanner.Scan() {
		if char := strings.ToLower(string(scanner.Text()[0]))[0]; char >= 97 && char <= 121 {
			wordCount[string(char)]++
		}
		wordCount[strings.ToLower(scanner.Text())]++
	}
	check(scanner.Err())
	fmt.Printf("Sorted by values: %v\n", sortMapByValue(wordCount))
	fmt.Printf("Sorted by keys: %v\n", sortMapByKeys(wordCount))

}
