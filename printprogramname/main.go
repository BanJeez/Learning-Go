package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	aba := argument[0][2:]

	for _, x := range aba {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}
