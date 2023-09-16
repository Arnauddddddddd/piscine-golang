package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var nombre int = 0
	if baseFrom == "01" {
		tab_nbr := []byte(nbr)
		calcul := 1
		for i := len(tab_nbr) - 1; i >= 0; i-- {
			if tab_nbr[i] == '1' {
				nombre += calcul
			}
			calcul *= 2
		}
	}
	return nbr
}
