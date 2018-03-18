package rectangle

import "math"

// Each package can have init() for initialization tasks
func init() {
    println("rectangle package initialized")
}

// Exported functions are capitalized
func Area(leng, wid float64) float64 {
    return leng * wid
}

func Diagonal(leng, wid float64) float64 {
    return math.Sqrt((leng * leng) + (wid * wid))
}
