package piscine

func NRune(s string, n int) rune {
	sentence := []rune(s)

	if len(sentence) >= n && n > 0 {
		return sentence[n-1]
	}
	return 0
}

// var Nrune rune
//x := []rune(s)
//if n > len(s)|| n < 0 {
//	return 0
//} else {
//	return x[n-1]
