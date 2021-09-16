package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	// aba := argument[1]

	for ai := 1; ai <= len(argument)-1; ai++ {
		for _, x := range argument[ai] {
			z01.PrintRune(rune(x))
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
