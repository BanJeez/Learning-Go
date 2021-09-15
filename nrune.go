package piscine

func NRune(s string, n int) rune {
	// var Nrune rune
	x := []rune(s)
	if n > len(s)-1 || n < 0 {
		return 0
	} else {
		return x[n-1]
	}
}
