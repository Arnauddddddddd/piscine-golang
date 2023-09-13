package piscine

func IterativeFactorial(nb int) int {
	entier := 1
	liste_multi := []int{}
	for nb != 0 {
		liste_multi = append(liste_multi, nb)
		nb--
	}
	for i := 0; i < len(liste_multi); i++ {
		entier *= liste_multi[i]
	}
	return entier
}
