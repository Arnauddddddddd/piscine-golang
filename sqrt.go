package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	p := 0
	for i := 0; nb > i; i++ {
		if nb >= i*(i-1) && nb <= i*i {
			if nb == i*i {
				p = i
			} else {
				return 0
			}
			break
		}
	}
	return p
}
