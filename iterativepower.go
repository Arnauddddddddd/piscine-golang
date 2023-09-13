package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || nb >= 99 || power < 0 {
		return 0
	}
	if power > 1 {
		return nb * IterativePower(nb, power - 1)
	} else {
		return nb
	}
}
