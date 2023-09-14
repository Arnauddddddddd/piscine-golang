package piscine

func TrimAtoi(s string) int {
	a := []byte(s)
	tab := []byte{}
	tab2 := []byte{}
	txt := ""
	for i := range a {
		if (rune(a[i]) >= 48 && rune(a[i]) <= 57) || rune(a[i]) == 45 {
			tab = append(tab, a[i])
		}
	}
	if len(tab) == 0 {
		return 0
	}
	if tab[0] == 45 {
		tab = tab[1:]
		tab2 = append(tab2, '-')
	}
	for i := range tab {
		if tab[i] != 45 {
			tab2 = append(tab2, tab[i])
		}
	}
	for i := range tab2 {
		txt += string(tab2[i])
	}
	return Atoi(txt)
}
