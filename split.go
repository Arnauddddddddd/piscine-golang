package piscine

func Split(s, sep string) []string {
	tab_indice := []int{}
	tab := []byte(s)
	sepa := []byte(sep)
	var compteur int = 0
	for i := 0; i < (len(tab) - len(sepa)); i++ {
		compteur = 0
		for j := 0; j < len(sepa); j++ {
			if tab[i+j] == sepa[j] {
				compteur += 1
			}
		}
		if compteur == len(sep) {
			tab_indice = append(tab_indice, i)
		}
	}
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab_indice); j++ {
			for k := 0; k < len(sepa); k++ {
				tab[tab_indice[j]+k] = ' '
			}
		}
	}
	txt := ""
	for i := 0; i < len(tab); i++ {
		txt += string(tab[i])
	}
	return SplitWhiteSpaces(txt)
}
