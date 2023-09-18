package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	tab := []rune(os.Args[0])
	for i := range tab {
		if i > 1 {
			z01.PrintRune(tab[i])
		}
	}
	z01.PrintRune('\n')
}
