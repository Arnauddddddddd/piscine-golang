package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	liste_int := []rune{}
	entier := 0
	up := false
	if len(args) >= 1 {
		if args[0] == "--upper" {
			up = true
			args = args[1:]
		}
		for i := 0; i < len(args); i++ {
			entier = 0
			for j := 0; j < len(args[i]); j++ {
				entier = 10*entier + int(args[i][j]) - 48
			}
			liste_int = append(liste_int, rune(entier))
		}
		for k := 0; k < len(liste_int); k++ {
			if !up {
				if (liste_int[k]+96 >= 65 && rune(liste_int[k]+96) <= 90) || (liste_int[k]+96 >= 97 && rune(liste_int[k]+96) <= 122) {
					z01.PrintRune(liste_int[k] + 96)
				} else {
					z01.PrintRune(32)
				}
			} else {
				if (liste_int[k]+64 >= 65 && rune(liste_int[k]+64) <= 90) || (liste_int[k]+64 >= 97 && rune(liste_int[k]+64) <= 122) {
					z01.PrintRune(liste_int[k] + 64)
				} else {
					z01.PrintRune(32)
				}
			}
		}
		z01.PrintRune('\n')
	}
}
