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
	for _, caractere := range s {
		entier = int(caractere - 48) + 10*entier
		if caractere < 48 || caractere > 57 {
			return 0
		}
	}
	return signe * entier
}
