package slices_count

import (
    _ "fmt"
    "testing"
)

/* -- stringLength function -- */
// In:  array of strings
// Out: minLength, maxLength
var stringLengths = []struct {
    input   []string
    wantMin int
    wantMax int
}{
    {[]string{"hell", "world"}, 4, 5},
    {[]string{"", ""}, 0, 0},
    {[]string{"aaa", "bbb", "c"}, 1, 3},
    {[]string{"四月份平民", ""}, 0, 5},
}

func TestStringLengths(t *testing.T) {
    for _, tt := range stringLengths {
        gotMin, gotMax := minMax(tt.input)

        if gotMin != tt.wantMin {
            t.Errorf("minMax(%v): gotMin: %d, wantMin: %d", tt.input, gotMin, tt.wantMin)
        }
        if gotMax != tt.wantMax {
            t.Errorf("minMax(%v): gotMax: %d, wantMAx: %d", tt.input, gotMax, tt.wantMax)
        }
    }
}

var in1 = []string{"Hello", "World"}
var want1 = [][]string{{"Hello", "World"}}

var in2 = []string{"aa", "bb", "ccc", "ddd"}
var want2 = [][]string{{"aa", "bb"}, {"ccc", "ddd"}}

var in3 = []string{""}
var want3 = [][]string{{""}}

var in4 = []string{"", "a", "bb", "cc", "ddd"}
var want4 = [][]string{{""}, {"a"}, {"bb", "cc"}, {"ddd"}}

var in5 = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
    "Emma", "Isabella", "Emily", "Madison",
    "Ava", "Olivia", "Sophia", "Abigail",
    "Elizabeth", "Chloe", "Samantha",
    "Addison", "Natalie", "Mia", "Alexis"}

// [Ava Mia] [Evan Neil Adam Matt Emma] [Emily Chloe] [Martin Olivia Sophia Alexis]
// [Katrina Madison Abigail Addison Natalie] [Isabella Samantha] [Elizabeth]
var want5 = [][]string{
    {"Ava", "Mia"},
    {"Evan", "Neil", "Adam", "Matt", "Emma"},
    {"Emily", "Chloe"},
    {"Martin", "Olivia", "Sophia", "Alexis"},
    {"Katrina", "Madison", "Abigail", "Addison", "Natalie"},
    {"Isabella", "Samantha"},
    {"Elizabeth"},
}

/* -- organize function -- */
var stringSlices = []struct {
    input []string
    want  [][]string
}{
    {in1, want1},
    {in2, want2},
    {in3, want3},
    {in4, want4},
    {in5, want5},
}

func TestOrganize(t *testing.T) {
    for _, tt := range stringSlices {
        got := Organize(tt.input)
        //fmt.Printf("In: %v, got: %v, want: %v\n", tt.input, got, tt.want)

        for i := range got {
            for j := range got[i] {
                if got[i][j] != tt.want[i][j] {
                    t.Errorf("organize(%v): got %v, want %v", tt.input, got, tt.want)
                }
            }
        }
    }
}
