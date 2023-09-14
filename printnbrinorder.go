package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	liste_nombre := []int{}
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for n > 0 {
			liste_nombre = append(liste_nombre, n%10)
			n = n / 10
		}
		for n < 0 {
			liste_nombre = append(liste_nombre, -(n % 10))
			n = n / 10
		}
		SortIntegerTable(liste_nombre)
		for i := range liste_nombre {
			z01.PrintRune(rune(liste_nombre[i] + 48))
		}
	}
}
