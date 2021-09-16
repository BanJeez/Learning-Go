package piscine

func ToUpper(s string) string {
	sentence := []rune(s)
	for index, x := range sentence {
		if x >= 'a' && x <= 'z' {
			sentence[index] = x - 32
		}
	}
	return string(sentence)
}
