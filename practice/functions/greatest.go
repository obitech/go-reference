package functions

// Finds greatest number in list of numbers
func greatest(nums ...int) int {
	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
