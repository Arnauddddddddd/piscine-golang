package piscine

func BasicJoin(elems []string) string {
	txt := ""
	for i := range elems {
		txt += elems[i]
	}
	return txt
}
