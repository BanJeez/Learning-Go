package piscine

// import "github.com/01-edu/z01"

func StrRev(s string) string {
	x := []rune{}

	for i := len(s) - 1; i >= 0; i-- {
		// fmt.Println(i)
		x = append(x, rune(s[i]))
	}
	return string(x)
}
