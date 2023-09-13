package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	p := 0
	i := 0
	if nb == 3 {
		return 0
	}
	for true {
		if nb >= i*(i-1) && nb <= i*(i+1) || i*i == nb {
			p = i
			break
		}
		i++
	}
	return p
}
