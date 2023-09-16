package piscine

func Index(s string, toFind string) int {
	s_tab := []byte(s)
	tofind_tab := []byte(toFind)
	var index int = -1
	var compteur int
	for i := 0; i < len(s_tab)-len(tofind_tab); i++ {
		for j := 0; j < len(tofind_tab); j++ {
			if s_tab[i+j] == tofind_tab[j] {
				compteur++
			}
		}
		if compteur == len(tofind_tab) {
			return i
		}
		compteur = 0
	}
	return index
}
