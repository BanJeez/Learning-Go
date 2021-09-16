package piscine

func IsNumeric(s string) bool {
	for _, x := range s {
		if x <= 47 || x >= 58 {
			return false
		}
	}
	return true
}
