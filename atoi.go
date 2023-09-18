package piscine

func Atoi(s string) int {
	entier := 0
	signe := 1
	if len(s) == 0 {
		return 0
	}
	if s[0] == '+' {
		signe = 1
		s = s[1:]
	}
	if s[0] == '-' {
		signe = -1
		s = s[1:]
	}
	
	for i := 0; i < len(s); i++ {
		if rune(s[i]) < 48 || rune(s[i]) > 57 {
			return 0
		}
		entier = 10 * entier
		entier += int(s[i] - 48)
	}
	return signe * entier
}
