package piscine

func Split(s, sep string) []string {
	test := []string{}
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
	for i := len(s) - 1; i > 0; i-- {
		for j := len(tab_indice) - 1; j > 0; j-- {
			s = s[:tab_indice[j]-1] + s[(tab_indice[j]):]
		}
	}
	txt := ""
	for i := 0; i < len(tab); i++ {
		txt += string(tab[i])
	}
	return test
}
