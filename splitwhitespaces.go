package piscine

func SplitWhiteSpaces(s string) []string {
	for i := 1; i < len(s); i++ {
		if string(s[i-1]) == " " && string(s[i]) == " " {
			s = string(s[:i-1] + s[i:])
			i--
		}
	}
	tab_str := []string{}
	tab := []byte(s)
	var txt string = ""
	for i := 0; i < len(tab); i++ {
		txt = ""
		if tab[i] == ' ' {
			for j := 0; j < i; j++ {
				txt = txt + string(tab[j])
			}
			tab_str = append(tab_str, txt)
			tab = tab[i+1:]
			i = 0
		}
	}
	txt = ""
	for i := range tab {
		txt = txt + string(tab[i])
	}
	tab_str = append(tab_str, txt)
	return tab_str
}
