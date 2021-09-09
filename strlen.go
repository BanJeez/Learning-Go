package piscine

func StrLen(s string) int {
	length := 0
	for i, c := range s {
		i = +i
		c = +c
		length += 1
	}

	return length
}
