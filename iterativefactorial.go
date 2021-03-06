package piscine

func IterativeFactorial(nb int) int {
	x := 1
	if nb > 30 {
		return 0
	} else if nb < 0 {
		return 0
	}

	for i := 1; i <= nb; i++ {
		x *= i
	}
	if x <= 0 {
		return 0
	} else if x == 1 {
		return 1
	}

	return x
}
