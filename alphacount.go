package piscine

func AlphaCount(s string) int {
	count := 0
	result := []rune(s)

	for i := 0; i < len(result); i++ {
		if (result[i] >= 'a' && result[i] <= 'z') || (result[i] >= 'A' && result[i] <= 'Z') {
			count++
		}
	}
	return count
}
