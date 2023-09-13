package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb < 0 || nb >= 99 {
		return 0
	}
	if nb < 2 {
		return nb
	} else {
		return nb * IterativeFactorial(nb-1)
	}
}
