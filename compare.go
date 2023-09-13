package piscine

func Compare(a, b string) int {
	valeur := 0
	if a == b {
		return valeur
	}
	if a < b {
		return -1
	} else {
		return 1
	}
}
