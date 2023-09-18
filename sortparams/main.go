package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	table := os.Args[1:]
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			z01.PrintRune(rune(table[i][j]))
		}
		z01.PrintRune('\n')
	}
}
