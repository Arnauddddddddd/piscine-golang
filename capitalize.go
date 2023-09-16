package piscine

func Capitalize(s string) string {
	s_tab := []rune(s)
	if (s_tab[0] >= 97 && s_tab[0] <= 122) {
		s_tab[0] -= 32
	}
	for i := 1; i < len(s_tab); i++ {
		if ((s_tab[i-1] >= 65 && s_tab[i-1] <= 90) || (s_tab[i-1] >= 97 && s_tab[i-1] <= 122) || (s_tab[i-1] >= 48 && s_tab[i-1] <= 57)) {
			if (s_tab[i] >= 65 && s_tab[i] <= 90) {
				s_tab[i] += 32
			}
		} else {
			if (s_tab[i] >= 97 && s_tab[i] <= 122) {
				s_tab[i] -= 32
			}
		}
	}
	txt := ""
	for i := 0; i < len(s_tab); i++ {
		txt += string(s_tab[i])
	}
	return txt
}
