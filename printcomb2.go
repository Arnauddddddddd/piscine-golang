package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	listeChiffre := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					if (i != k || j != l) && j < i {
						z01.PrintRune(listeChiffre[i])
						z01.PrintRune(listeChiffre[j])
						z01.PrintRune(' ')
						z01.PrintRune(listeChiffre[k])
						z01.PrintRune(listeChiffre[l])
						if i != 9 || j != 8 || k != 9 || l != 9 {
							z01.PrintRune(rune(44))
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune(rune(10))
}
