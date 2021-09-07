package main

import "github.com/01-edu/z01"

func main() {
	var aRune string = "abcdefghijklmnopqrstuvwxyz"

	// z01.PrintRune(rune(aRune[i]))

	for i := 0; i <= 25; i++ {
		// fmt.Println(i)
		z01.PrintRune(rune(aRune[i]))
	}
}
