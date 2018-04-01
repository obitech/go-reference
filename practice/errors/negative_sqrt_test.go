package negative_sqrt

import (
	"fmt"
	"math"
	"testing"
)

var numsPass = []struct {
	input float64
	want  float64
}{
	{0.0, math.Sqrt(0)},
	{1.0, math.Sqrt(1.0)},
	{3.14, math.Sqrt(3.14)},
	{1000000000, math.Sqrt(1000000000)},
}

func testSqrtPass(t *testing.T) {
	for _, tt := range numsPass {
		got, err := MySqrt(tt.input)

		if got != tt.want || err != nil {
			t.Errorf("Sqrt(%f): got '%f', want '%f', err '%v'", tt.input, got, tt.want, err)
		}
	}
}

func fmtError(t *testing.T, got float64) string {
	t.Helper()
	return fmt.Sprintf("cannot Sqrt negative number: %f", got)
}

// var numsFail = []struct {
//     input float64
//     want  float64
//     err   string
// }{
//     {0, 0, ""},
//     {-2, 0, fmtError(t, -2)},
// }

// func testSqrtFail(t *testing.T) {
//     for _, tt := range numsFail {
//         got, err := MySqrt(tt.input)

//         if got != want || err == nil {
//             t.Errorf("Sqrt(%f): got '%f', want '%f', err '%s'", tt.input, got, tt.want, err)
//         }
//     }
// }
