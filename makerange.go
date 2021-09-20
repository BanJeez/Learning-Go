package piscine

func MakeRange(min, max int) []int {
	ii := min
	var B []int

	if max < min || max == min {
		return B
	}
	A := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		A[i] = ii
		ii++
	}
	return A
}
