package piscine

func SplitWhiteSpaces(s string) []string {
	slicedRune := []rune(s)

	var answer []string

	var result string

	for i := 0; i < len(slicedRune); i++ {
		if slicedRune[i] == ' ' || slicedRune[i] == '\n' || slicedRune[i] == '\t' {
			if result != "" {
				answer = append(answer, result)
				result = ""
			}
		} else {
			result = result + string(slicedRune[i])
		}
		if i == len(slicedRune)-1 {
			answer = append(answer, result)
		}
	}
	return answer
}
