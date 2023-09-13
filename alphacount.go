package piscine

func AlphaCount(s string) int {
	tab := []rune(s)
	compteur := 0
	for i:= range tab {
		if (tab[i] >= 65 && tab[i] <= 90) || (tab[i] >= 97 && tab[i] <= 122) {
			compteur += 1
		}
	}
	return compteur
}
