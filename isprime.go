package piscine

func IsPrime(nb int) bool {
	if nb % 2 == 0 || nb % 3 == 0 {
		return false
	} else {
		return true
	}
}
