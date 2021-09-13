package piscine

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	x := z01.PrintRune(rune(s[0]))
	return rune(x)
}
