package piscine

func ConcatParams(args []string) string {
	R := ""
	// var X []string
	for _, x := range args {
		R += x + "\n"
	}
	return string(R)
}
