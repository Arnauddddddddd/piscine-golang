package piscine 

func IsPrintable(s string) bool {
	for i := range s {
		if rune(s[i]) < 33 || rune(s[i]) > 126 {
			return false
		}
	}
	return true
}
