package piscine

func IsAlpha(s string) bool {
	for _, x := range s {
		if x <= 47 || x >= 58 && x <= 64 || x > 90 && x < 97 || x > 122 {
			return false
		}
	}
	return true
}
