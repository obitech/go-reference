package functions

import (
	"testing"
)

var argsStruct = []struct {
	input    int
	wantInt  int
	wantBool bool
}{
	{1, 0, false},
	{2, 1, true},
	{0, 0, true},
	{24, 12, true},
}

func TestHalf(t *testing.T) {
	for _, tt := range argsStruct {
		gotInt, gotBool := half(tt.input)

		if gotInt != tt.wantInt {
			t.Errorf("half(%d): gotInt: %d, wantInt: %d", tt.input, gotInt, tt.wantInt)
		}

		if gotBool != tt.wantBool {
			t.Errorf("hal(%d): gotBool: %v, wantBool: %v", tt.input, gotBool, tt.wantBool)
		}
	}
}
