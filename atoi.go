package piscine

func Atoi(s string) int {
	entier := 0
	signe := 1
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
	for _, caractere := range s {
		entier = 10 * entier
		entier += int(caractere - 48)
		if caractere < 48 || caractere > 57 {
			return 0
		}
	}
	return signe * entier
}
