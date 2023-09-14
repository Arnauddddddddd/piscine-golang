package piscine

func BasicJoin(elems []string) string {
	txt := ""
	for i := range elems {
		txt += elems[i]
		if !(i == len(elems)-1) {
			txt += " "
		}
	}
	return txt
}
