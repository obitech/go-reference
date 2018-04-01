package functions

// Function takes an int, returns arg divided by two and bool if arg was even
func half(arg int) (int, bool) {
	return arg / 2, arg%2 == 0
}
