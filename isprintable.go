package piscine

func IsPrintable(s string) bool {
	for _, x := range s {
		if x <= 31 {
			return false
		}
	}
	return true
}
