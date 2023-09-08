package piscine

func Strlen(s string) {
	compte := 0
	for i := 0; i < len(s); i++ {
		compte++
	}
	return compte
}
