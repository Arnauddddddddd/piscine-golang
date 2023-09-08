package piscine

func StrLen(s string) int {
	compte := 0
	for i := 0; i < len(s)-1; i++ {
		compte++
	}
	return compte
}
