package piscine

func IterativeFactorial(nb int) int {
	x := 1
	for i := 1; i <= nb+1; i++ {
		x = i * i
	}
	return x - 1
}
