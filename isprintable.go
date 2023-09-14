package piscine

func IsPrintable(s string) bool {
	for i := range s {
		if rune(s[i]) < 32 || rune(s[i]) > 126 {
			return false
		}
	}
	return true
}
