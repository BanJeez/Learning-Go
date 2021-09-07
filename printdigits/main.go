package main

import "github.com/01-edu/z01"

func main() {
	var aRune string = "00123456789"

	// z01.PrintRune(rune(aRune[i]))

	for i := 0; i <= 9; {
		i++
		// fmt.Println(i)
		z01.PrintRune(rune(aRune[i]))

	}
	z01.PrintRune('\n')
}
