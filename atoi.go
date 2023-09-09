package piscine

func Atoi(s string) int {
	ListeChiffre := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ListeChiffreRune := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var nombre int = 0
	var puissance int = 1
	for element := len(s) - 1; element >= 0; element-- {
		for chiffre := 0; chiffre < len(ListeChiffreRune); chiffre++ {
			if s[element] == byte(ListeChiffreRune[chiffre]) {
				if chiffre != '0' {
					nombre = puissance*ListeChiffre[chiffre] + nombre
				}
				puissance = puissance * 10
			}
			if (rune(s[element]) < 48 || rune(s[element]) > 57) && rune(s[element]) != 45 && rune(s[element]) != 43 {
				return 0
			}
		}
	}
	for i := 1; i < len(s); i++ {
		if rune(s[i]) == 45 || rune(s[i]) == 43 {
			return 0
		}
	}
	if len(s) > 0 {
		if rune(s[0]) == 45 {
			nombre -= nombre * 2
		}
	}
	return nombre
}
