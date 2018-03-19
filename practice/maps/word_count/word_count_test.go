package word_count

import (
    _ "fmt"
    "testing"
)

var words = []struct {
    inWords       []string
    inCoins       int
    wantWordCount map[string]int
    wantCoins     int
}{
    {[]string{"Hello"}, 3, map[string]int{
        "Hello": 2,
    }, 1},
    {[]string{"Jim", "John", "Jack"}, 3, map[string]int{
        "Jim":  1,
        "John": 1,
        "Jack": 1,
    }, 0},
    {[]string{"Jim", "John", "Jack", "Jones"}, 3, map[string]int{
        "Jim":   1,
        "John":  1,
        "Jack":  1,
        "Jones": 0,
    }, 0},
    {[]string{"Jim", "John", "Jack"}, 0, map[string]int{
        "Jim":  0,
        "John": 0,
        "Jack": 0,
    }, 0},
    {[]string{"Xxx", "TRSWDXD", "LLlLkkMn"}, 10, map[string]int{
        "Xxx":      0,
        "TRSWDAD":  0,
        "LLlLkkMn": 0,
    }, 10},
    {[]string{"aaaeeeiiiooo", "John", "Jack"}, 12, map[string]int{
        "aaaeeeiiiooo": 10,
        "John":         1,
        "Jack":         1,
    }, 0},
    {[]string{""}, 5, map[string]int{}, 5},
    {[]string{"", "Peer"}, 5, map[string]int{
        "":     0,
        "Peer": 2,
    }, 3},
    {[]string{"四月份平民"}, 5, map[string]int{"四月份平民": 0}, 5},
}

func mapsEqual(t *testing.T, m1 map[string]int, m2 map[string]int) bool {
    t.Helper()
    for k, v := range m1 {
        if m2[k] != v {
            return false
        }
    }
    return true
}

func TestWordCount(t *testing.T) {
    for _, tt := range words {
        gotWordCount, gotCoins := WordCount(tt.inWords, tt.inCoins)

        //fmt.Printf("In: %v, %d\ngotWordCount: %v, wantWordCount: %v\ngotCoins: %d, wantCoins: %d\n", tt.inWords, tt.inCoins, gotWordCount, tt.wantWordCount, gotCoins, tt.wantCoins)

        if !mapsEqual(t, gotWordCount, tt.wantWordCount) {
            t.Errorf("WordCount(%v, %d): gotWordCount: %v, wantWordCount: %v", tt.inWords, tt.inCoins, gotWordCount, tt.wantWordCount)
        }

        if gotCoins != tt.wantCoins {
            t.Errorf("WordCount(%v, %d): gotCoins: %d, wantCoins: %d", tt.inWords, tt.inCoins, gotCoins, tt.wantCoins)
        }
    }
}
