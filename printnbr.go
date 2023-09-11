package piscine

func PrintNbr(n int) {
	n2 := n
	liste_nombre_retire := []int{0}
	texte := ""
	signe := 1
	multiplicateur := 1
	if n < 0 {
		n = n - n -n
		n2 = n2 - n2 - n2
		signe = -1
	}
	var entier int
	var nombre int
	puissance := 0
	for i := 0; n >= 10 ; i++{
		n = n / 10
		puissance++
	}
	for j := 0; puissance != 0; j++{
		multiplicateur *= 10
		puissance--
	}
	for multiplicateur != 0 {
		nombre = n2
		for _, element := range liste_nombre_retire {
			nombre = nombre - element
		}
		entier = nombre / multiplicateur
		liste_nombre_retire = append(liste_nombre_retire, entier*multiplicateur)
		multiplicateur /= 10
		texte += string(entier+48)
	}
	if signe == -1 {
		texte = "-" + texte
	}
}
