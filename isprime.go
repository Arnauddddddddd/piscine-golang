package piscine

func IsPrime(nb int) bool {
	if nb < 0 {
		return false
	}
	if nb == 1 {
		return false
	}
	if nb == 2 || nb == 3 || nb == 5 || nb == 7 {
		return true
	}
	for i := 11; i < nb || i > 500; i++ {
		if !(nb%2 == 0 || nb%3 == 0 || nb%5 == 0 || nb%7 == 0) {
			return false
		}
	}
	return true
}
