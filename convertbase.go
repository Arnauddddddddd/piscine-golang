package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var n int = 0
	if baseFrom == "01" {
		tab_nbr := []byte(nbr)
		calcul := 1
		for i := len(tab_nbr) - 1; i >= 0; i-- {
			if tab_nbr[i] == '1' {
				n += calcul
			}
			calcul *= 2
		}
		liste_nombre := []int{}
		txt := ""
		if n == 0 {
			return "0"
		} else {
			for n > 0 {
				liste_nombre = append(liste_nombre, n%10)
				n = n / 10
			}
			for n < 0 {
				liste_nombre = append(liste_nombre, -(n % 10))
				n = n / 10
			}
			for i := len(liste_nombre) - 1; i >= 0; i-- {
				txt += string((rune(liste_nombre[i] + 48)))
			}
			return txt
		}
	}
	return nbr
}
