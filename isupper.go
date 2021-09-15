package piscine

func IsUpper(s string) bool {
	for _, x := range s {
		if x < 65 || x > 90 {
			return false
		}
	}
	return true
}
