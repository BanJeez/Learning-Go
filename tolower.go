package piscine

func ToLower(s string) string {
	sentence := []rune(s)
	for index, x := range sentence {
		if x >= 'A' && x <= 'Z' {
			sentence[index] = x + 32
		}
	}
	return string(sentence)
}
