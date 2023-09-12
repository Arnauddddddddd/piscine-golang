package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	liste_nombre := []int{}
	if n < 0 {
		n = -n
		z01.PrintRune(45)
	}
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			liste_nombre = append(liste_nombre, n%10)
			n = n / 10
		}
		for i := len(liste_nombre) - 1; i >= 0; i-- {
			z01.PrintRune(rune(liste_nombre[i] + 48))
		}
	}
}
