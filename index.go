package piscine

func Index(s string, toFind string) int {
	sentence := []rune(s)
	find := len([]rune(toFind))

	if len(sentence) < find {
		return -1
	}
	for i := 0; i < len(sentence); i++ {
		if len(sentence[i:]) < len([]rune(toFind)) {
			return -1
		} else {
			if s[i:i+find] == toFind {
				return i
			}
		}
	}
	return -1
}
