package piscine

func IsUpper(s string) bool {
	for i := range s {
		if rune(s[i]) < 65 || rune(s[i]) > 90 {
			return false
		}
	}
	return true
}
