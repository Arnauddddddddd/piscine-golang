package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	tab := []rune(s)
	if len(tab) == 0 {
		return 0
	}
	if n > len(tab) {
		return 0
	}
	return tab[n-1]
}
