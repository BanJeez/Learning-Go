package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a:= 48 ; a < 56 ; a++{
		for b:= a + 1 ; b < 57 ; b++{
			for c:= b + 1 ; c <58 ; c++{
				if a ==55 && b == 56 && k == 57{
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					z01.PrintRune('\n')
				}else {
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			} 
		}
	}
}
