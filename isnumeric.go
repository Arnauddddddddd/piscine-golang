package piscine

func IsNumeric(s string) bool {
	for i := range s {
		if rune(s[i]) < 48 || rune(s[i]) > 57 {
			return false
		}
	}
	return true
}
