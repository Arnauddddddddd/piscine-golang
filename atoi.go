package piscine

func Atoi(s string) int {
	entier := 0
	signe := 1
	s2 := []rune(s)
	if len(s) == 0 {
		return 0
	}
	if s2[0] == '+' {
		signe = 1
		s2 = s2[1:]
	}
	if s2[0] == '-' {
		signe = -1
		s2 = s2[1:]
	}
	for _, i := range s2 {
		if i < 48 || i > 57 {
			return 0
		}
		entier = 10*entier + int(i-48)
	}
	return signe * entier
}
