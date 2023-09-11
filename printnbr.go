package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	liste_nombre_retire := []int{0}
	liste_rune := []rune{}
	signe := 1
	multiplicateur := 1
	if n < 0 {
		n = n - n*2
		signe = -1
	}
	n2 := n
	var entier int
	var nombre int
	puissance := 0
	for i := 0; n >= 10; i++ {
		n = n / 10
		puissance++
	}
	for j := 0; puissance != 0; j++ {
		multiplicateur *= 10
		puissance--
	}
	if signe == -1 {
		z01.PrintRune(45)
	}
	for multiplicateur != 0 {
		nombre = n2
		for _, element := range liste_nombre_retire {
			nombre = nombre - element
		}
		entier = nombre / multiplicateur
		liste_nombre_retire = append(liste_nombre_retire, entier*multiplicateur)
		multiplicateur /= 10
		liste_rune = append(liste_rune, rune(entier+48))
	}
	for _, k := range liste_rune {
		z01.PrintRune(k)
	}
	z01.PrintRune(' ')
}
