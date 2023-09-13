package piscine

func IterativePower(nb int, power int) int {
	if power > 1 {
		return nb * IterativePower(nb, power-1)
	} else {
		return nb
	}
}
