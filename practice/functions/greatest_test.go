package functions

import (
	"testing"
)

var numsStruct = []struct {
	input []int
	want  int
}{
	{[]int{1, 2, 3, 4}, 4},
	{[]int{}, 0},
	{[]int{-1, -2, -3, 500}, 500},
}

func TestGreatest(t *testing.T) {
	for _, tt := range numsStruct {
		got := greatest(tt.input...)

		if got != tt.want {
			t.Errorf("greatest(%d): got %d, want %d", tt.input, got, tt.want)
		}
	}
}
