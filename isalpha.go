package piscine

func IsAlpha(s string) bool {
	for _, x := range s {
		if x <= 47 || x >= 58 && x <= 64 || x > 90 && x < 97 || x > 'z' {
			return false
		}
	}
	return true
}
