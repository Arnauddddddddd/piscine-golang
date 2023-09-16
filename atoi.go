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
	if len(s) == 0 {
		return 0
	}
	s2 := []byte(s)
	for i := 0; i < len(s2); i++ {
		if rune(s2[i]) < 48 || rune(s2[i]) > 57 {
			return 0
		}
		entier = 10 * entier
		entier += int(s2[i] - 48)
	}
	return signe * entier
}
