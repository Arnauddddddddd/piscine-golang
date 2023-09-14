package piscine

func ToLower(s string) string {
	tab := []rune(s)
	for i := range tab {
		if tab[i] >= 65 && tab[i] <= 90 {
			tab[i] = rune(s[i] + 32)
		}
	}
	s2 := ""
	for i := 0; i < len(tab); i++ {
		s2 += string(tab[i])
	}
	return s2
}
