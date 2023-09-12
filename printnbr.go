package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	liste_nombre := []int{}
	var signe int
	multiplicateur := 1
	m := n
	var entier int
	if n < 0 {
		signe = -1
		m, n = -n, -n
	}
	if signe == -1 {
		z01.PrintRune(45)
	}
	if n == 0 {
		z01.PrintRune('0')
	} else {
		for m > 10 {
			m = m / 10
			multiplicateur *= 10
		}
		for n > 1 {
			entier = n / multiplicateur
			liste_nombre = append(liste_nombre, entier)
			n = n - entier*multiplicateur
			multiplicateur /= 10
		}
		for _, i := range liste_nombre {
			z01.PrintRune(rune(i + 48))
		}
	}
}
