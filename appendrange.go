package piscine

func AppendRange(min, max int) []int {
	var A []int

	for i := min - 1; i < max-1; i++ {
		A = append(A, i+1)
	}
	return A
}
