package piscine

func LastRune(s string) rune {
	var last rune
	for _, c := range s {
		last = c
	}
	return rune(last)
}
