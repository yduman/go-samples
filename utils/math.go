package utils

// Add multiple numbers together and return the sum of it
func Add(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
