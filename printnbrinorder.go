package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	liste_nombre := []int{}
	signe := 0
	if n < 0 {
		signe = -1
	}
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
		for i := 0; i <= len(liste_nombre) - 1; i++ {
			z01.PrintRune(rune(liste_nombre[i] + 48))
		}
		if signe == -1 {
			z01.PrintRune(45)
		}
	}
}
