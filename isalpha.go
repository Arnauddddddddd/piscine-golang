package piscine

func IsAlpha(s string) bool {
	for i := range s {
		if (rune(s[i]) < 97 || rune(s[i]) > 122) && (rune(s[i]) < 65 || rune(s[i]) > 90) && (rune(s[i]) < 48 || rune(s[i]) > 57) {
			return false
		}
	}
	return true
}
