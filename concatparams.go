package piscine

func ConcatParams(args []string) string {
	R := ""
	// var X []string
	for i, x := range args {
		if i < len(args)-1 {
			R += x + "\n"
		} else {
			R += x
		}
	}
	return string(R)
}
