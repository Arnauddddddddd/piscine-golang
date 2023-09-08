package piscine

func BasicAtoi2(s string) int {
	ListeChiffre := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ListeChiffreRune := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var nombre int = 0
	var puissance int = 1
	var compteur int = 0
	for element := len(s) - 1; element >= 0; element-- {
		for chiffre := 0; chiffre < len(ListeChiffreRune); chiffre++ {
			if compteur > 12 {
				return 0
			}
			if s[element] == byte(ListeChiffreRune[chiffre]) {
				if chiffre != '0' {
					nombre = puissance*ListeChiffre[chiffre] + nombre
				}
				compteur = 0
				puissance = puissance * 10
			}
			compteur = compteur + 1
		}
	}
	return nombre
}
