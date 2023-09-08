package piscine

import "github.com/01-edu/z01"

func Strlen(s string) {
	compte := 0
	for i := 0; i < len(s); i++ {
		compte++
	}
	z01.PrintRune(rune(compte))
}
