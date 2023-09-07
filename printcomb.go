package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	listeChiffre := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if i < j && j < k {
					z01.PrintRune(listeChiffre[i])
					z01.PrintRune(listeChiffre[j])
					z01.PrintRune(listeChiffre[k])
					if i == 7 && j == 8 && k == 9 {
						continue
					} else {
						z01.PrintRune(rune(44)) 
					}					
				}
				z01.PrintRune(' ')
			}
		}
	}
}
