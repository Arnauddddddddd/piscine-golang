package piscine

func Index(s string, toFind string) int {
	s_tab := []byte(s)
	tofind_tab := []byte(toFind)
	var index int
	var compteur int
	var test bool = false
	for i := 0; i < len(s_tab)-len(tofind_tab); i++ {
		for j := 0; j < len(tofind_tab); j++ {
			if s_tab[i+j] == tofind_tab[j] {
				compteur++
				test = true
			}
		}
		if compteur == len(tofind_tab) {
			index++
		}
		compteur = 0
	}
	if test == true {
		return index
	} else {
		return -1
	}
}
