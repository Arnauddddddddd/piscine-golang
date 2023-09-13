package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	tab := []rune(s)
	if n > len(tab) - 1 {
		return 0
	}
	return tab[n-1]
}
