package piscine

func Atoi(s string) int {
	entier := 0
	signe := 1
	s2 := []rune(s)
	if len(s2) <= 0 {
		return 0
	}
	if s2[0] == '+' {
		signe = 1
		s2 = s2[1:]
	}
	if len(s2) <= 0 {
		return 0
	}
	if s2[0] == '-' {
		signe = -1
		s2 = s2[1:]
	}
	if len(s2) <= 0 {
		return 0
	}
	for _, i := range s2 {
		if i < '0' || i > '9' {
			return 0
		}
		entier = 10*entier + int(i-'0')
	}
	if len(s2) <= 0 {
		return 0
	}
	return signe * entier
}
