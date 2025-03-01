package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	if power > 1 {
		return nb * IterativePower(nb, power-1)
	} else {
		return nb
	}
}
