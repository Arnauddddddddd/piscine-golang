package piscine

func ToUpper(s string) string {
	tab := []rune(s)
	for i := range tab {
		if tab[i] >= 97 && tab[i] <= 122 {
			tab[i] = rune(s[i] - 32)
		}
	}
	s2 := ""
	for i := 0; i < len(tab); i++ {
		s2 += string(tab[i])
	}
	return s2
}
