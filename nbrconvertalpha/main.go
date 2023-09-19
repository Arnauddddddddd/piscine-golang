package main

import (
	"os"

	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		z01.PrintRune(rune(piscine.Atoi(args[i]))+96)
	}
}
