package piscine

func Join(strs []string, sep string) string {
	txt := ""
	for i := range strs {
		txt += strs[i]
		if !(i == len(strs)-1) {
			txt += sep
		}
	}
	return txt
}