package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for i, c := range s {
		i = +i
		z01.PrintRune(rune(c))
	}
}
